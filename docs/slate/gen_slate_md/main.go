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
	"path/filepath"
	"os"
	"github.com/go-openapi/spec"
	"regexp"
)

var buildDir = flag.String("build-dir", "", "destination for generated files")
var sourceDir = flag.String("source-dir", "", "source for swagger files")
var templateDir = flag.String("template-dir", "", "")
var conceptsList = flag.String("concepts", "v1.Pod,v1.Container,v1beta1.Deployment,v1.ObjectMeta", "comma-separated list of top level concepts")

func main() {
	flag.Parse()

	var specs []*loads.Document = []*loads.Document{}
	err := filepath.Walk(*sourceDir, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext != ".json" {
			return nil
		}
		d, err := loads.JSONSpec(path)
		if err != nil {
			return fmt.Errorf("Could not load json file %s as api-sec: %v", path, err)
		}
		specs = append(specs, d)
		return nil
	})
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
		os.Exit(1)
	}

	topLevelConcepts := map[string]*template.ConceptTemplateParams{}
	for _, concept := range strings.Split(*conceptsList, ",") {
		topLevelConcepts[concept] = &template.ConceptTemplateParams{}
	}

	// Process concepts
	conceptList := map[string]spec.Schema{}
	for _, d := range specs {
		for k, v := range d.Spec().Definitions {
			conceptList[k] = v
		}
	}

	operationList := map[string]spec.PathItem{}

	// Process endpoints
	for _, d := range specs {
		for path, pathItem := range d.Spec().Paths.Paths {
			var concept *template.ConceptTemplateParams = nil
			for tlcName, ctp := range topLevelConcepts {
				re := regexp.MustCompile(strings.Replace(strings.ToLower(tlcName), ".", "/(.*/)?", 1) + "(s)?[$/]")

				if re.Match([]byte(path)) {
					concept = ctp
					break
				}
			}
			if pathItem.Get != nil {
				operationList["g" + path] = pathItem
			}
			if pathItem.Post != nil {
				operationList["p" + path] = pathItem
			}
			if pathItem.Delete != nil {
				operationList["d" + path] = pathItem
			}
			if pathItem.Put != nil {
				operationList["u" + path] = pathItem
			}
			if pathItem.Patch != nil {
				operationList["a" + path] = pathItem
			}
			if concept == nil {
				continue
			}
			fmt.Printf("%s %s\n", concept, path)
			pathParams := []spec.Parameter{}
			for _, p := range pathItem.Parameters {
				pathParams = append(pathParams, p)
			}
			if pathItem.Get != nil {
				PrintOp("GET", pathItem.Get, pathParams, concept)
			}
			if pathItem.Post != nil {
				PrintOp("POST", pathItem.Post, pathParams, concept)
			}
			if pathItem.Delete != nil {
				PrintOp("DELETE", pathItem.Delete, pathParams, concept)
			}
			if pathItem.Put != nil {
				PrintOp("PUT", pathItem.Put, pathParams, concept)
			}
			if pathItem.Patch != nil {
				PrintOp("PATCH", pathItem.Patch, pathParams, concept)
			}
		}
	}
	fmt.Printf("Length of Concepts: %d\n", len(conceptList))
	fmt.Printf("Length of Operations: %d\n", len(operationList))

	 //fmt.Printf("Result %+s\n", d)

	//ct, err := template.New(*templateDir)
	//if err != nil { panic(err) }
	//err = ct.Execute(os.Stdout, GetDeployment())
	//if err != nil { panic(err) }
}

func PrintOp(t string, op *spec.Operation, pathParams []spec.Parameter, c *template.ConceptTemplateParams) {
	if op != nil {
		fmt.Printf("\t%s: %s\n", t, op.Description)
		fmt.Printf("\t\tPathParameters:\n")
		for _, p := range pathParams {
			if p.Type != "" {
				fmt.Printf("\t\t-%s %s [%+v]\n", p.Name, p.Description, p.Type)
			} else {
				fmt.Printf("\t\t-%s %s [%+v]\n", p.Name, p.Description, p.Schema.SchemaProps.Ref.GetPointer())
			}
		}
		fmt.Printf("\t\tQueryParameters:\n")
		for _, p := range op.Parameters {
			if p.Type != "" {
				fmt.Printf("\t\t-%s %s [%+v]\n", p.Name, p.Description, p.Type)
			} else {
				fmt.Printf("\t\t-%s %s [%+v]\n", p.Name, p.Description, p.Schema.SchemaProps.Ref.GetPointer())
			}
		}
		if r, f := op.Responses.ResponsesProps.StatusCodeResponses[200]; f && len(r.Schema.SchemaProps.Ref.GetPointer().String()) > 0 {
			fmt.Printf("\t\tResponses:\n")
			fmt.Printf("\t\t-%v\n", r.Schema.SchemaProps.Ref.GetPointer())
		}
	}
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