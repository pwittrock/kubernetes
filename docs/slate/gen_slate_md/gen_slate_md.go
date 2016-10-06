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
	"fmt"
	"flag"

	"github.com/go-openapi/loads"
)

var dst = flag.String("dst", "", "destination for generated files")
var src = flag.String("src", "", "source for swagger files")

func main() {
	flag.Parse()

	d, err := loads.JSONSpec(*src)
	if err != nil {
		fmt.Printf("Could not load json %v\n", err)
		return
	}

	for k, v := range d.Spec().Paths.Paths {
		fmt.Printf("%s\n", k)
		if v.Get != nil {
			fmt.Printf("\tGET: %s -> %s\n", k, v.Get.Description)
		}
		if v.Post != nil {
			fmt.Printf("\tPOST: %s -> %s\n", k, v.Post.Description)
		}
		if v.Delete != nil {
			fmt.Printf("\tDELETE: %s -> %s\n", k, v.Delete.Description)
		}
		if v.Put != nil {
			fmt.Printf("\tPUT: %s -> %s\n", k, v.Put.Description)
		}
		if v.Patch != nil {
			fmt.Printf("\tPATCH: %s -> %s\n", k, v.Patch.Description)
		}
	}
	//
	//fmt.Printf("Result %+s\n", d)
}