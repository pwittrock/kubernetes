/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/apis/extensions/v1beta1"
	"regexp"
)

var tagWhitelist = flag.String("tag-whitelist", "", "Only push updates for image tags matching this regular expression.  Default to everything.")
var controllerAnnotation = flag.String("controller-annotation", "replicatingperfection.net/push-image", "Push images only to controllers with this annotation key.")
var namespace = flag.String("namespace", "", "Namespace to run against.")
var pushLabelPrefix = flag.String("push-label-prefix", "auto-pushed-image-", "Update this label on the PodTemplate for each auto pushed image.")

var _ DeploymentUpdater = &deploymentUpdaterImpl{}

type deploymentUpdaterImpl struct {
	clients *kubernetes.Clientset
	tagRegex *regexp.Regexp
}

// DeploymentUpdater performs updates to deployments
type DeploymentUpdater interface {
	// Update Deployments IFF
	// - The request contains a secret param matching --web-hook-secret
	// - The Deployment as an annotation key matching --controller-annotation
	// - The result of Unmarshalling the request body has a non-empty repository name
	// - The image tag matches --tag-whitelist
	// - The Deployment has a container matching the container name/respository
	// Deployments will be updated by setting an annotation in the PodTemplate and updating the container image
	// label to match the label from the webhook
	// TODO: Only update the container if the webhook build time is greater than the last update buildtime
	// repoName: name of image repository e.g. username/imagename
	// tag: tag of pushed image e.g. latest
	// pushedAt: unix timestamp of when image was pushed with this tag to trigger the callback
	UpdateDeployments(repoName, tag string, pushedAt int)

	// TagMatches returns true if the tag for the image is whitelisted for autopushing
	TagMatches(tag string) bool
}

// NewDeploymentUpdater creates a new DeploymentUpdater
func NewDeploymentUpdater(clientset *kubernetes.Clientset) DeploymentUpdater {
	return &deploymentUpdaterImpl{clientset, regexp.MustCompile(*tagWhitelist)}
}

// UpdateDeployments Updates Deployment PodTemplate images by applying the new tag
func (h *deploymentUpdaterImpl) UpdateDeployments(repoName, tag string, pushedAt int) {
	// Find Deployments that need to be updated and push the new image
	deployments, err := h.clients.Extensions().Deployments(*namespace).List(api.ListOptions{})
	if err != nil {
		fmt.Printf("Failed to process request due to error listing Deployments: %v\n", err)
		return
	}
	for _, d := range deployments.Items {
		if h.shouldPushImageToDeployment(repoName, pushedAt, d) {
			if err := h.updateDeploymentContainerImages(repoName, tag, pushedAt, d); err != nil {
				fmt.Printf("Failed to update Deployment %s image with %s:%s due to error %v.\n",
					d.Name, repoName, tag, err)
			}
		}
	}
}

// updateDeploymentContainerImages updates the container images for a single deployment
func (h *deploymentUpdaterImpl) updateDeploymentContainerImages(cbRepoName, tag string, pushedAt int, d v1beta1.Deployment) error {
	template := d.Spec.Template

	// Delete any labels for images that are no longer used by any containers
	keepImageLabels := map[string]bool{}
	for i := range d.Spec.Template.Spec.Containers {
		// Parse out the repo name
		c := &d.Spec.Template.Spec.Containers[i]
		parts := strings.Split(c.Image, ":")
		if len(parts) < 1 {
			return fmt.Errorf("Could not parse image name from %s.\n", c.Image)
		}
		// Index the label
		keepImageLabels[getPushLabel(parts[0])] = true
	}
	for k := range template.Labels {
		// Any labels for container images that are used anymore should be removed
		if isPushLabel(k) && !keepImageLabels[k] {
			delete(template.Labels, k)
		}
	}

	// Update container images
	for i := range d.Spec.Template.Spec.Containers {
		// Parse out the repo name
		c := &d.Spec.Template.Spec.Containers[i]
		parts := strings.Split(c.Image, ":")
		if len(parts) < 1 {
			return fmt.Errorf("Could not parse image name from %s.\n", c.Image)
		}
		currRepoName := parts[0]

		// Check if it matches the repo for this container
		if cbRepoName != currRepoName {
			continue
		}

		existingTag := ""
		if len(parts) > 1 {
			existingTag = parts[1]
		}
		if !h.TagMatches(existingTag) {
			// Tag whitelist does not match existing container tag.  Don't update.
			continue
		}

		imageName := ""
		if len(tag) > 0 {
			// Add the image tag if specified
			imageName = cbRepoName + ":" + tag
		} else {
			imageName = cbRepoName
		}

		// Set the new image
		c.Image = imageName

		// Update the label on the pod with the push time
		template.Labels[getPushLabel(cbRepoName)] = fmt.Sprintf("%d", pushedAt)
	}

	_, err := h.clients.Extensions().Deployments(d.Namespace).Update(&d)
	return err
}

// TagMatches returns true if the tag for the image is whitelisted for autopushing
func (h *deploymentUpdaterImpl) TagMatches(tag string) bool {
	// Check that the tag matches the whitelist of tags to update
	if len(*tagWhitelist) > 0 {
		if !h.tagRegex.Match([]byte(tag)) {
			return false
		}
	}
	return true
}

// GetPushLabel returns the Label key applied for each updated image
func getPushLabel(repoName string) string {
	return *pushLabelPrefix + repoName
}

// IsPushLabel returns true if the label key is for pushed images
func isPushLabel(label string) bool {
	return strings.HasPrefix(label, *pushLabelPrefix)
}

// ShouldPushImageToDeployment returns true if the deployment requires and image update
func (h *deploymentUpdaterImpl) shouldPushImageToDeployment(cbRepoName string, newPushTime int, d v1beta1.Deployment) bool {
	// Check for a matching annotation on the deployment
	hasAnnotation := false
	for k := range d.Annotations {
		if k == *controllerAnnotation {
			hasAnnotation = true
			break
		}
	}

	// Filter deployments without annotation
	if !hasAnnotation {
		return false
	}
	fmt.Printf("Deployment %s found annotation %s.\n", d.Name, *controllerAnnotation)

	template := d.Spec.Template
	if p, f := template.Labels[getPushLabel(cbRepoName)]; f {
		lastPushTime, err := strconv.Atoi(p)
		if err != nil {
			fmt.Printf("Error parsing time for label %s on deployment %s (%s).\n", getPushLabel(cbRepoName), d.Name, p)
			return false
		}
		if newPushTime < lastPushTime {
			fmt.Printf("Skipping image %s on deployment %s.  New push time %d < old push time %d.\n",
				cbRepoName, d.Name, newPushTime, lastPushTime)
			return false
		}
	}

	// Filter deployments matching container
	hasImage := false
	for _, c := range d.Spec.Template.Spec.Containers {
		// Parse out the image from the tag
		parts := strings.Split(c.Image, ":")
		if len(parts) < 1 {
			fmt.Printf("Could not parse image name from %s.\n", c.Image)
			return false
		}
		currRepoName := parts[0]

		existingTag := ""
		if len(parts) > 1 {
			existingTag = parts[1]
		}
		if !h.TagMatches(existingTag) {
			// Tag whitelist does not match existing container tag.  Don't update.
			fmt.Printf("Skipping image: [%s] on deployment: [%s] because tag: [%s] does not match whitelist: [%s].\n", currRepoName, d.Name, existingTag, *tagWhitelist)
			continue
		}

		// Compare container image against the image from the callback
		if cbRepoName == currRepoName {
			fmt.Printf("Image %s matches Deployment %s image %s.\n", cbRepoName, d.Name, currRepoName)
			hasImage = true
			break
		} else {
			fmt.Printf("Skipping image: [%s] on deployment: [%s] because image does not match pushed repo: [%s].\n", currRepoName, d.Name, cbRepoName)
		}
	}
	return hasImage
}
