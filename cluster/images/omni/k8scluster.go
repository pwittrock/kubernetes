/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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
//	_ "github.com/fsouza/go-dockerclient"

//	_ "github.com/golang/glog"
	"fmt"
)

//var client *docker.Client

func init() {
	flag.Parse()
//	client, err := docker.NewClientFromEnv()
//	if err != nil {
//		glog.Fatal("Could not create Docker client %v", err)
//	}
}

func main() {
	cmd := flag.Args()[0]
	switch cmd {
	case "connect":
	case "delete":
	case "start":
	case "stop":
	}

	fmt.Println(cmd)
}

func start() {
}