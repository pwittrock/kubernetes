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

package template

import (
	"text/template"
)

type ConceptTemplate struct {
	*template.Template
}

func New(path string) (*ConceptTemplate, error) {
	t, err := template.New("concept.template").ParseFiles(path)
	if err != nil { return nil, err }
	temp := &ConceptTemplate{t}
	return temp, nil
}

type ConceptTemplateParams struct {
	Name string
	Fields []ConceptField
	Operations []Operation
}

type ConceptField struct {
	Name string
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