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

package dockerhub

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pwittrock/container-image-pusher/controllers"

	"k8s.io/client-go/1.4/kubernetes"
)

var expectedSecret = flag.String("web-hook-secret", "12345",
	"Secret used to authenticate requests as coming from the dockerhub.  The webhook callback provide this as a query param in the form /?secret=123345")

// Interface type checking
var _ http.Handler = &WebHookCallback{}

// WebHookCallback implements handling of DockerHub webhook callbacks by updating Deployment container images with image updates
type WebHookCallback struct {
	controllers.DeploymentUpdater
}

// NewWebhookCallback is a Constructor
func NewWebhookCallback(clients *kubernetes.Clientset) (*WebHookCallback, error) {
	return &WebHookCallback{controllers.NewDeploymentUpdater(clients)}, nil
}

// ServeHTTP handles a webhook callback by updating the PodTemplate container images
func (h *WebHookCallback) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Handling request %+s\n", req.RequestURI)

	body, err := h.ParseCallback(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("Processing request.  Parsed Body: %+v Values %v URI: %s.\n", body, req.URL.Query(), req.RequestURI)

	repoName := body.Repository.RepoName
	tag := body.PushData.Tag
	pushedAt := body.PushData.PushedAt

	// Verify the tag on the image matches the whitelist
	if !h.TagMatches(tag) {
		return
	}

	// Update Deployments
	h.UpdateDeployments(repoName, tag, pushedAt)
}

// ParseCallback parses the webhook callback request into a struct
// - Verify the secret if specified
// - Verify the result is non-tempty
func (h *WebHookCallback) ParseCallback(req *http.Request) (*CallbackBody, error) {

	// Read out the request body into a struct
	bodyTxt, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read request body %v\n", req)
	}

	// Authenticate the request is from our webhook callback by checking for the secret value
	actualSecret := req.URL.Query().Get("secret")
	if actualSecret != *expectedSecret {
		return nil, fmt.Errorf("Unauthorized request.  Secret does not match %s (%s): %+v\n", *expectedSecret, actualSecret, req)
	}

	body := &CallbackBody{}
	err = json.Unmarshal(bodyTxt, body)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse body %s\n", bodyTxt)
	}
	if len(body.Repository.RepoName) < 1 {
		return nil, fmt.Errorf("Unable to process DockerHub callback %+v.  Missing RepoName.\n", body)
	}
	return body, nil
}
