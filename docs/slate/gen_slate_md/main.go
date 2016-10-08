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
	//"os"

	"github.com/go-openapi/loads"

	"k8s.io/kubernetes/docs/slate/gen_slate_md/slate"
	"strings"
	"errors"
)

var dst = flag.String("dst", "", "destination for generated files")
var src = flag.String("src", "", "source for swagger files")
var templates = flag.String("templates", "", "")
var conceptsList = flag.String("concepts", "v1.PodSpec,v1.Container,v1beta1.Deployment,v1.ObjectMeta", "comma-separated list of top level concepts")

func main() {
	flag.Parse()

	d, err := loads.JSONSpec(*src)
	if err != nil {
		fmt.Printf("Could not load json %v\n", err)
		return
	}

	topLevelConcepts := map[string]string{}
	for _, value := range strings.Split(*conceptsList, ",") {
		topLevelConcepts[value] = ""
		fmt.Printf("%s\n", value)
	}

	for k, _ := range d.Spec().Definitions {
		if _, f := topLevelConcepts[k]; f {
			fmt.Printf("Def: %s\n", k)
		}
	}

	//for k, v := range d.Spec().Paths.Paths {
	//	fmt.Printf("%s\n", k)
	//	if v.Get != nil {
	//		fmt.Printf("\tGET: %s -> %s\n", k, v.Get.Description)
	//		for _, p := range v.Get.Parameters {
	//			fmt.Printf("\t\t%s Param %s %s %+v\n", k, p.Name, p.Description, p.Type)
	//		}
	//	}
	//	if v.Post != nil {
	//		fmt.Printf("\tPOST: %s -> %s\n", k, v.Post.Description)
	//		for _, p := range v.Post.Parameters {
	//			fmt.Printf("\t\t%s Param %s %s %s %+v\n", k, p.Name, p.Description, p.Type, p.Schema.SchemaProps.Ref.GetPointer())
	//		}
	//	}
	//	if v.Delete != nil {
	//		fmt.Printf("\tDELETE: %s -> %s\n", k, v.Delete.Description)
	//	}
	//	if v.Put != nil {
	//		fmt.Printf("\tPUT: %s -> %s\n", k, v.Put.Description)
	//	}
	//	if v.Patch != nil {
	//		fmt.Printf("\tPATCH: %s -> %s\n", k, v.Patch.Description)
	//	}
	//}

	 //fmt.Printf("Result %+s\n", d)

	//ct, err := template.New(*templates)
	//if err != nil { panic(err) }
	//err = ct.Execute(os.Stdout, GetDeployment())
	//if err != nil { panic(err) }
}

type Concepts []string

func (c *Concepts) String() string {
	return strings.Join(*c, ",")
}

func (c *Concepts) Set(value string) error {
	if len(*c) > 0 {
		return errors.New("concepts flag already set")
	}
	*c = strings.Split(value, ",")
	return nil
}

func GetDeployment() template.ConceptTemplateParams {
	return template.ConceptTemplateParams{
		Name: "Deployment",
		Fields: []template.ConceptField{
			{"apiVersion", "string", "Must be set to 'extensions/v1beta1'"},
			{"kind", "string", "Must be set to 'extensions/v1beta1'"},
		},
		Operations: []template.Operation{
			{
				"Create",
				"Create a new Deployment",
				"POST https://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/{namespace}/deployments",
				[]template.Example{
					{"shell", "Some Shell Request", "Some Shell Result"},
					{"ruby", "Some Ruby Request", "Some Ruby Result"}},
				[]template.PathParam{
					{"namespace", "string", "Object name and auth scope, such as for teams and projects"},
				},
				[]template.QueryParam{
					{"pretty", "bool", "If 'true', then the output is pretty printed."},
					{"labelSelector", "string", "A selector to restrict the list of returned objects by their labels. Defaults to everything."},
				},
				[]template.RequestBody{
					{"body", "[Deployment](#Deployment)", "Deployment resource object"},
				},
			},
			{
				"Delete",
				"Delete a Deployment",
				"Delete https://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}",
				[]template.Example{
					{"shell", "Some Shell Request", "Some Shell Result"},
					{"ruby", "Some Ruby Request", "Some Ruby Result"}},
				[]template.PathParam{
					{"namespace", "string", "Object name and auth scope, such as for teams and projects"},
					{"name", "string", "name of the Deployment"},
				},
				[]template.QueryParam{
					{"pretty", "bool", "If 'true', then the output is pretty printed."},
					{"labelSelector", "string", "A selector to restrict the list of returned objects by their labels. Defaults to everything."},
				},
				[]template.RequestBody{
				},
			},
		},
	}

}