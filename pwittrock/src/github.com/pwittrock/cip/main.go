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

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/rest"

	"github.com/pwittrock/container-image-pusher/dockerhub"
)

// Use encryption by default
var port = flag.Int("port", 443, "Port to listen for connections.")
var certPath = flag.String("cert-path", "/tls/tls.crt", "Path to certificate file for https." +
	"Must be set if --key-path is set.  If unset, don't use encrypted connections.")
var keyPath = flag.String("key-path", "/tls/tls.key", "Path to key file for https.  " +
	"Must be set if --cert-path is set.  If unset, don't use encrypted connections.")

func main() {
	flag.Parse()

	clients := GetConfig()
	h, err := dockerhub.NewWebhookCallback(clients)
	if err != nil {
		fmt.Printf("Failed to initialize %v.\n", err)
	}

	if (len(*certPath) == 0 || len(*keyPath) == 0) && len(*certPath) != len(*keyPath) {
		fmt.Printf("Both --cert-path and --key-path must be specified if either is specified.\n")
		os.Exit(2)
	}

	fmt.Printf("Listening on :%d.\n", *port)
	portStr := fmt.Sprintf(":%d", *port)
	if len(*certPath) > 0 {
		err = http.ListenAndServeTLS(portStr, *certPath, *keyPath, h)
	} else {
		fmt.Printf("Warning, serving on unencrypted end point.\n")
		err = http.ListenAndServe(portStr, h)
	}
	if err != nil {
		fmt.Printf("Failed to listen for connections. %v\n", err)
		os.Exit(3)
	}
}

// GetConfig returns a set of kubernetes api clients
func GetConfig() *kubernetes.Clientset {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clients, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clients
}
