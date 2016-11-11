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

package api

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/go-openapi/loads"

	"github.com/pwittrock/slate-tools/lib"
)

func NewConfig(yamlFile, openApiFile string) *Config {
	config := loadYamlConfig(yamlFile)
	specs := LoadOpenApiSpec(openApiFile)

	// Initialize all of the operations
	config.Definitions = GetDefinitions(specs)

	// Mark resources in the ToC as such
	markToc := func(resource *Resource, definition *Definition) {
		definition.InToc = true
		resource.Definition = definition
	}
	config.VisitResourcesInToc(config.Definitions, markToc)

	// Get the map of operations appearing in the open-api spec keyed by id
	config.InitOperations(specs)
	config.CleanUp()

	return config
}

// GetOperations returns all Operations found in the Documents
func (config *Config) InitOperations(specs []*loads.Document) {
	o := Operations{}
	VisitOperations(specs, func(operation Operation) {
		//fmt.Printf("Operation: %s\n", operation.ID)
		o[operation.ID] = &operation
	})
	config.Operations = o

	config.mapOperationsToDefinitions()
	config.Definitions.initializeOperationParameters(config.Operations)
}

// CleanUp sorts and dedups fields
func (c *Config) CleanUp() {
	for _, d := range c.Definitions.GetAllDefinitions() {
		sort.Sort(d.AppearsIn)
		sort.Sort(d.Fields)
		dedup := SortDefinitionsByName{}
		last := ""
		for _, i := range d.AppearsIn {
			if i.Name == last {
				continue
			}
			last = i.Name
			dedup = append(dedup, i)
		}
		d.AppearsIn = dedup
	}
}

// loadYamlConfig reads the config yaml file into a struct
func loadYamlConfig(yamlFile string) *Config {
	config := &Config{}
	if len(*lib.YamlFile) < 1 {
		fmt.Printf("Must specify --yaml-file.\n")
		os.Exit(2)
	}

	contents, err := ioutil.ReadFile(*lib.YamlFile)
	if err != nil {
		fmt.Printf("Failed to read yaml file %s: %v", yamlFile, err)
		os.Exit(2)
	}

	err = yaml.Unmarshal(contents, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return config
}

// initOpExample reads the example config for each operation and sets it
func (config *Config) initOpExample(o *Operation) {
	path := o.Type.Name + ".yaml"
	path = filepath.Join(config.ExampleLocation, o.Definition.Name, path)
	path = strings.Replace(path, " ", "_", -1)
	path = strings.ToLower(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, &o.ExampleConfig)
	if err != nil {
		panic(fmt.Sprintf("Could not Unmarshal yaml: %s\n", content))
	}
}

func getOperationId(match string, group ApiGroup, version ApiVersion, kind string) string {
	// Substitute the api definition group-version-kind into the operation template and look for a match
	v, k := doScaleIdHack(string(version), kind, match)
	match = strings.Replace(match, "${group}", string(group), -1)
	match = strings.Replace(match, "${version}", v, -1)
	match = strings.Replace(match, "${resource}", k, -1)
	return match
}

func (config *Config) setOperation(match string, group ApiGroup, namespaceRep string,
	ot *OperationType, oc *OperationCategory,  definition *Definition) {
	key := strings.Replace(match, "(Namespaced)?", namespaceRep, -1)
	if o, found := config.Operations[key]; found && o.Definition == nil {
		// Each definition should be in exactly 1 group
		if len(definition.Group) > 0 && group != definition.Group {
			panic(fmt.Sprintf(
				"Multiple Groups found for Resource %v %v %v\n",
				definition.Name, definition.Group, group))
		}
		definition.Group = group
		o.Type = *ot
		o.Definition = definition
		oc.Operations = append(oc.Operations, o)
		config.initOpExample(o)
	}
}

// mapOperationsToDefinitions adds operations to the definitions they operate
// This is done by - for each definition - look at all potentially matching operations from operation categories
func (config *Config) mapOperationsToDefinitions() {
	// Look for matching operations for each definition
	for _, definition := range config.Definitions.GetAllDefinitions() {
		// Iterate through categories
		for i := range config.OperationCategories {
			oc := config.OperationCategories[i]
			// Iterate through possible operation matches
			for j := range oc.OperationTypes {
				// Iterate through possible api groups since we don't know the api group of the definition
				for _, group := range config.ApiGroups {
					ot := oc.OperationTypes[j]
					operationId := getOperationId(ot.Match, group, definition.Version, definition.Name)
					// Look for a matching operation and set on the definition if found
					config.setOperation(operationId, group, "Namespaced", &ot, &oc, definition)
					config.setOperation(operationId, group, "", &ot, &oc, definition)
				}
			}

			// If we found operations for this category, add the category to the definition
			if len(oc.Operations) > 0 {
				definition.OperationCategories = append(definition.OperationCategories, &oc)
			}
		}
	}
}

func doScaleIdHack(version, name, match string) (string, string) {
	// Hack to get around ids
	if strings.HasSuffix(match, "${resource}Scale") && name != "Scale" {
		// Scale names don't generate properly
		name = strings.ToLower(name) + "s"
		out := []rune(name)
		out[0] = unicode.ToUpper(out[0])
		name = string(out)
	}
	out := []rune(version)
	out[0] = unicode.ToUpper(out[0])
	version = string(out)

	return version, name
}
