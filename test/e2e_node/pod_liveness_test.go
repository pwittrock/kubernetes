/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package e2e_node

import (
	"fmt"
	"time"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/util/intstr"
	"k8s.io/kubernetes/test/e2e/framework"

	. "github.com/onsi/ginkgo"
)

var _ = framework.KubeDescribe("Pods", func() {
	f := framework.NewDefaultFramework("pods")
	It("should be restarted with a docker exec \"cat /tmp/health\" liveness probe [Conformance]", func() {
		runLivenessTest(f.Client, f.Namespace.Name, &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name:   "liveness-exec",
				Labels: map[string]string{"test": "liveness"},
			},
			Spec: api.PodSpec{
				Containers: []api.Container{
					{
						Name:    "liveness",
						Image:   "gcr.io/google_containers/busybox:1.24",
						Command: []string{"/bin/sh", "-c", "echo ok >/tmp/health; sleep 10; rm -rf /tmp/health; sleep 600"},
						LivenessProbe: &api.Probe{
							Handler: api.Handler{
								Exec: &api.ExecAction{
									Command: []string{"cat", "/tmp/health"},
								},
							},
							InitialDelaySeconds: 15,
							FailureThreshold:    1,
						},
					},
				},
			},
		}, 1, defaultObservationTimeout)
	})

	It("should *not* be restarted with a docker exec \"cat /tmp/health\" liveness probe [Conformance]", func() {
		runLivenessTest(f.Client, f.Namespace.Name, &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name:   "liveness-exec",
				Labels: map[string]string{"test": "liveness"},
			},
			Spec: api.PodSpec{
				Containers: []api.Container{
					{
						Name:    "liveness",
						Image:   "gcr.io/google_containers/busybox:1.24",
						Command: []string{"/bin/sh", "-c", "echo ok >/tmp/health; sleep 600"},
						LivenessProbe: &api.Probe{
							Handler: api.Handler{
								Exec: &api.ExecAction{
									Command: []string{"cat", "/tmp/health"},
								},
							},
							InitialDelaySeconds: 15,
							FailureThreshold:    1,
						},
					},
				},
			},
		}, 0, defaultObservationTimeout)
	})

	It("should be restarted with a /healthz http liveness probe [Conformance]", func() {
		runLivenessTest(f.Client, f.Namespace.Name, &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name:   "liveness-http",
				Labels: map[string]string{"test": "liveness"},
			},
			Spec: api.PodSpec{
				Containers: []api.Container{
					{
						Name:    "liveness",
						Image:   "gcr.io/google_containers/liveness:e2e",
						Command: []string{"/server"},
						LivenessProbe: &api.Probe{
							Handler: api.Handler{
								HTTPGet: &api.HTTPGetAction{
									Path: "/healthz",
									Port: intstr.FromInt(8080),
								},
							},
							InitialDelaySeconds: 15,
							FailureThreshold:    1,
						},
					},
				},
			},
		}, 1, defaultObservationTimeout)
	})

	// Slow by design (5 min)
	It("should have monotonically increasing restart count [Conformance] [Slow]", func() {
		runLivenessTest(f.Client, f.Namespace.Name, &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name:   "liveness-http",
				Labels: map[string]string{"test": "liveness"},
			},
			Spec: api.PodSpec{
				Containers: []api.Container{
					{
						Name:    "liveness",
						Image:   "gcr.io/google_containers/liveness:e2e",
						Command: []string{"/server"},
						LivenessProbe: &api.Probe{
							Handler: api.Handler{
								HTTPGet: &api.HTTPGetAction{
									Path: "/healthz",
									Port: intstr.FromInt(8080),
								},
							},
							InitialDelaySeconds: 5,
							FailureThreshold:    1,
						},
					},
				},
			},
		}, 5, time.Minute*5)
	})

	It("should *not* be restarted with a /healthz http liveness probe [Conformance]", func() {
		runLivenessTest(f.Client, f.Namespace.Name, &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name:   "liveness-http",
				Labels: map[string]string{"test": "liveness"},
			},
			Spec: api.PodSpec{
				Containers: []api.Container{
					{
						Name:  "liveness",
						Image: "gcr.io/google_containers/nettest:1.7",
						// These args are garbage but the image will exit if they're not there
						// we just care about /read serving a 200, which it always does.
						Args: []string{
							"-service=liveness-http",
							"-peers=1",
							"-namespace=" + f.Namespace.Name},
						Ports: []api.ContainerPort{{ContainerPort: 8080}},
						LivenessProbe: &api.Probe{
							Handler: api.Handler{
								HTTPGet: &api.HTTPGetAction{
									Path: "/read",
									Port: intstr.FromInt(8080),
								},
							},
							InitialDelaySeconds: 15,
							TimeoutSeconds:      10,
							FailureThreshold:    1,
						},
					},
				},
			},
		}, 0, defaultObservationTimeout)
	})
})

// TODO: Dedup this with the e2e/pods.go test
const (
	defaultObservationTimeout = time.Minute * 2
)

// TODO: Dedup this with the e2e/pods.go test
func runLivenessTest(c *client.Client, ns string, podDescr *api.Pod, expectNumRestarts int, timeout time.Duration) {
	By(fmt.Sprintf("Creating pod %s in namespace %s", podDescr.Name, ns))
	_, err := c.Pods(ns).Create(podDescr)
	framework.ExpectNoError(err, fmt.Sprintf("creating pod %s", podDescr.Name))

	// At the end of the test, clean up by removing the pod.
	defer func() {
		By("deleting the pod")
		c.Pods(ns).Delete(podDescr.Name, api.NewDeleteOptions(0))
	}()

	// Wait until the pod is not pending. (Here we need to check for something other than
	// 'Pending' other than checking for 'Running', since when failures occur, we go to
	// 'Terminated' which can cause indefinite blocking.)
	framework.ExpectNoError(framework.WaitForPodNotPending(c, ns, podDescr.Name),
		fmt.Sprintf("starting pod %s in namespace %s", podDescr.Name, ns))
	framework.Logf("Started pod %s in namespace %s", podDescr.Name, ns)

	// Check the pod's current state and verify that restartCount is present.
	By("checking the pod's current state and verifying that restartCount is present")
	pod, err := c.Pods(ns).Get(podDescr.Name)
	framework.ExpectNoError(err, fmt.Sprintf("getting pod %s in namespace %s", podDescr.Name, ns))
	initialRestartCount := api.GetExistingContainerStatus(pod.Status.ContainerStatuses, "liveness").RestartCount
	framework.Logf("Initial restart count of pod %s is %d", podDescr.Name, initialRestartCount)

	// Wait for the restart state to be as desired.
	deadline := time.Now().Add(timeout)
	lastRestartCount := initialRestartCount
	observedRestarts := 0
	for start := time.Now(); time.Now().Before(deadline); time.Sleep(2 * time.Second) {
		pod, err = c.Pods(ns).Get(podDescr.Name)
		framework.ExpectNoError(err, fmt.Sprintf("getting pod %s", podDescr.Name))
		restartCount := api.GetExistingContainerStatus(pod.Status.ContainerStatuses, "liveness").RestartCount
		if restartCount != lastRestartCount {
			framework.Logf("Restart count of pod %s/%s is now %d (%v elapsed)",
				ns, podDescr.Name, restartCount, time.Since(start))
			if restartCount < lastRestartCount {
				framework.Failf("Restart count should increment monotonically: restart cont of pod %s/%s changed from %d to %d",
					ns, podDescr.Name, lastRestartCount, restartCount)
			}
		}
		observedRestarts = restartCount - initialRestartCount
		if expectNumRestarts > 0 && observedRestarts >= expectNumRestarts {
			// Stop if we have observed more than expectNumRestarts restarts.
			break
		}
		lastRestartCount = restartCount
	}

	// If we expected 0 restarts, fail if observed any restart.
	// If we expected n restarts (n > 0), fail if we observed < n restarts.
	if (expectNumRestarts == 0 && observedRestarts > 0) || (expectNumRestarts > 0 &&
	observedRestarts < expectNumRestarts) {
		framework.Failf("pod %s/%s - expected number of restarts: %t, found restarts: %t",
			ns, podDescr.Name, expectNumRestarts, observedRestarts)
	}
}