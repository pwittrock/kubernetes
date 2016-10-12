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

package slate

import (
	"text/template"
)

type ConceptTemplate struct {
	*template.Template
}

func New(path string) (*ConceptTemplate, error) {
	t, err := template.New("concept.template").ParseFiles(path + "/concept.template")
	if err != nil { return nil, err }
	temp := &ConceptTemplate{t}
	return temp, nil
}

type IndexTemplateParams struct {
	ConceptIncludes []string
	DefinitionIncludes []string
}

type Operations []Operation
func (a Operations) Len() int           { return len(a) }
func (a Operations) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Operations) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type ConceptTemplateParams struct {
	Name string
	Fields []ConceptField
	Operations Operations
	SubConcepts []ConceptTemplateParams
}

type ConceptField struct {
	Name string
	Link string
	Schema string
	Description string
}

type Operation struct {
	Name string
	Description string
	HttpRequest string
	Examples []Example
	PathParams []PathParam
	QueryParams []QueryParam
	RequestBody []RequestBody
	HttpResponses HttpResponses
}

type HttpResponses []HttpResponse
func (a HttpResponses) Len() int           { return len(a) }
func (a HttpResponses) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a HttpResponses) Less(i, j int) bool {
	return a[i].Code < a[j].Code
}

type HttpResponse struct {
	Code string
	Description string
	Schema string
}

type Example struct {
	Lang string
	Request string
	Response string
}

type PathParam struct {
	Name string
	Schema string
	Description string
}

type QueryParam struct {
	Name string
	Schema string
	Description string
}

type RequestBody struct {
	Name string
	Schema string
	Description string
}