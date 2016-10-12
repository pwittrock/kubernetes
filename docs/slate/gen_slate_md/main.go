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

//$ go run gen_slate_md/main.go
//$ docker run -v $(pwd)/source:/slate/source -v $(pwd)/build:/slate/build pwittrock/slate
package main

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"k8s.io/kubernetes/docs/slate/gen_slate_md/slate"
)

var buildDir = flag.String("build-dir", "source/", "Destination for generated files.  Should be the slate/source/ directory.")
var sourceDir = flag.String("source-dir", "../../api/openapi-spec/", "Source for swagger files.  Should be the api/open-api/ directory.")
var templateDir = flag.String("template-dir", "gen_slate_md/slate/", "Location of templates.  Should be slate/gen_slate_md/slate/ directory.")
var pkgDir = flag.String("pkg-dir", "../../pkg", "Directory containing Kubernetes register.go files for discovering top level objects.")

func main() {
	flag.Parse()

	// Parse top level concepts from register.go files
	tlc := GetTopLevelConcepts()

	// Load the open-api specs
	docs := LoadDocs()

	// Index all of the resource definitions
	definitions := CreateDefinitions(docs, tlc)

	// Add supported operations (e.g. CRUD) to each definitions
	AddOperations(docs, definitions, tlc)

	// Write the index file importing each of the top level concept files
	WriteIndexFile(tlc, definitions)

	// Write each concept file imported by the index file
	WriteConceptFiles(tlc, definitions)

	// Write each definition file imported by the index file
	WriteDefinitionFiles(definitions)

	 //fmt.Printf("Result %+s\n", d)

	//ct, err := template.New(*templateDir)
	//if err != nil { panic(err) }
	//err = ct.Execute(os.Stdout, GetDeployment())
	//if err != nil { panic(err) }
}

type TopLevelConcepts []string

func (a TopLevelConcepts) Len() int           { return len(a) }
func (a TopLevelConcepts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TopLevelConcepts) Less(i, j int) bool {
	ai := strings.Split(a[i], ".")
	aj := strings.Split(a[j], ".")
	if len(ai) < 2 {
		fmt.Printf("Missing %s\n", a[i])
		return false
	}
	if len(aj) < 2 {
		fmt.Printf("Missing %s\n", a[j])
		return true
	}
	if ai[1] == aj[1] {
		aiv := 0
		switch ai[0] {
		case "v1":
			aiv = 0
		case "v1beta1":
			aiv = 1
		case "v1alpha1":
			aiv = 2
		}
		ajv := 0
		switch aj[0] {
		case "v1":
			aiv = 0
		case "v1beta1":
			aiv = 1
		case "v1alpha1":
			aiv = 2
		}
		return aiv < ajv
	} else {
		return ai[1] < aj[1]
	}
}

func GetTopLevelConcepts() TopLevelConcepts {
	tl := TopLevelConcepts{}

	out, e := exec.Command("find", *pkgDir, "-name", "register.go", "-type", "f").CombinedOutput()
	if e != nil {
		fmt.Printf("Failed to find top level objects: %v %s", e, out)
		os.Exit(1)
	}

	for _, f := range strings.Split(string(out), "\n") {
		cmdLine := fmt.Sprintf("grep '&[A-Za-z]*{},' %s | sed 's/.*&//;s/{},//' | sort | uniq", f)
		out, e := exec.Command("bash", "-c", cmdLine).CombinedOutput()
		if e != nil {
			fmt.Printf("Failed to find top level objects: %v", e)
			os.Exit(1)
		}
		o := strings.TrimSpace(string(out))

		re := regexp.MustCompile("/(v1[0-9a-z]*)/register.go$")
		version := re.FindString(f)
		if len(version) == 0 {
			continue
		}

		version = strings.Replace(version, "/register.go", "", 1)
		version = strings.Replace(version, "/", "", 1)
		for _, a := range strings.Split(o, "\n") {
			a = version + "." + strings.TrimSpace(a)
			if len(a) == 0 {
				continue
			}
			tl = append(tl, a)
		}
	}
	tl = append(tl, "v1.ObjectMeta")
	sort.Sort(tl)
	pruned := TopLevelConcepts{}
	last := ""
	for _, i := range tl {
		j :=  strings.Split(i, ".")
		if j[1] == last {
			continue
		}
		last = j[1]
		pruned = append(pruned, i)
	}
	return pruned
}

func WriteIndexFile(tlc TopLevelConcepts, definitions map[string]*ConceptDefinition) {
	// Add concepts
	conceptIncludes := TopLevelConcepts{}
	for _, name := range tlc {
		if concept, found := definitions[name]; found {
			if len(concept.templateParams.Operations) > 0 {
				conceptIncludes = append(conceptIncludes, GetConceptImport(name))
			}
		}
	}

	// Add definitions
	orderedDefinitions := TopLevelConcepts{}
	for name, _ := range definitions {
		orderedDefinitions = append(orderedDefinitions, name)
	}
	sort.Sort(orderedDefinitions)
	definitionIncludes := TopLevelConcepts{}
	last := ""
	for _, name := range orderedDefinitions {
		if last == GetSimpleName(name) {
			continue
		}
		last = GetSimpleName(name)
		definitionIncludes = append(definitionIncludes, GetDefinitionImport(name))
	}

	// Open the index file
	f, err := os.Create(*buildDir + "/index.html.md")
	if err != nil {
		fmt.Printf("Failed to open index: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	// Create the template
	t, err := template.New("index.template").ParseFiles(*templateDir + "/index.template")
	if err != nil {
		fmt.Printf("Failed to parse template: %v", err)
		os.Exit(1)
	}

	// Write the instantiated template to the index file
	err = t.Execute(f, slate.IndexTemplateParams{
		ConceptIncludes: conceptIncludes,
		DefinitionIncludes: definitionIncludes,
	})
	if err != nil {
		fmt.Printf("Failed to parse template: %v", err)
		os.Exit(1)
	}
}

func WriteConceptFiles(tlc TopLevelConcepts, definitions map[string]*ConceptDefinition) {
	// Setup the template to be instantiated
	t, err := template.New("concept.template").ParseFiles(*templateDir + "/concept.template")
	if err != nil {
		fmt.Printf("Failed to parse template: %v", err)
		os.Exit(1)
	}

	for _, name := range tlc {
		if definition, found := definitions[name]; found {
			if len(definition.templateParams.Operations) > 0 {
				WriteConcept(t, &definition.templateParams, name)
			}
		} else {
			fmt.Printf("WARNING: No definition found for top level object %s\n", name)
		}
	}
}

func WriteDefinitionFiles(definitions map[string]*ConceptDefinition) {
	// Setup the template to be instantiated
	t, err := template.New("definition.template").ParseFiles(*templateDir + "/definition.template")
	if err != nil {
		fmt.Printf("Failed to parse template: %v", err)
		os.Exit(1)
	}

	for name, concept := range definitions {
		WriteDefinition(t, &concept.templateParams, name)
	}
}

func WriteConcept(t *template.Template, c *slate.ConceptTemplateParams, name string) {
	WriteTemplate(t, c, GetConceptFilePath(name))
}

func WriteDefinition(t *template.Template, c *slate.ConceptTemplateParams, name string) {
	WriteTemplate(t, c, GetDefinitionFilePath(name))
}

func WriteTemplate(t *template.Template, data interface{}, path string) {
	conceptFile, err := os.Create(path)
	defer conceptFile.Close()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
	t.Execute(conceptFile, data)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}

// Loads all of the open-api documents
func LoadDocs() []*loads.Document {
	docs := []*loads.Document{}
	err := filepath.Walk(*sourceDir, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext != ".json" {
			return nil
		}
		d, err := loads.JSONSpec(path)
		if err != nil {
			return fmt.Errorf("Could not load json file %s as api-sec: %v", path, err)
		}
		docs = append(docs, d)
		return nil
	})
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
	return docs
}

// Parse all of the concept definitions from the loaded documents
func CreateDefinitions(docs []*loads.Document, tlc TopLevelConcepts) map[string]*ConceptDefinition {
	definitions := map[string]*ConceptDefinition{}
	for _, doc := range docs {
		for name, definition := range doc.Spec().Definitions {
			definitions[name] = CreateConceptDefinition(name, definition)
		}
	}
	//AddConceptSubConcepts(definitions, tlc)
	return definitions
}

var nameRegex *regexp.Regexp = regexp.MustCompile("\\..*")

func GetSimpleName(s string) string {
	return strings.Replace(nameRegex.FindString(s), ".", "", 1)
}

func getImport(s string) string {
	return "generated_" + strings.ToLower(strings.Replace(s, ".", "_", 50))
}

func toFileName(s string) string {
	return fmt.Sprintf("%s/includes/_%s.md", *buildDir, s)
}

func GetDefinitionImport(s string) string {
	return fmt.Sprintf("%s_definition", getImport(s))
}

func GetDefinitionFilePath(s string) string {
	return toFileName(GetDefinitionImport(s))
}

func GetConceptImport(s string) string {
	return fmt.Sprintf("%s_concept", getImport(s))
}

func GetConceptFilePath(s string) string {
	return toFileName(GetConceptImport(s))
}

func CreateConceptDefinition(name string, s spec.Schema) *ConceptDefinition {
	fields := []slate.ConceptField{}
	for propertyName, propertyValue := range s.Properties {
		field :=  slate.ConceptField{
			Name: propertyName,
			Description: strings.Replace(propertyValue.Description, "\n", " ", 10000),
			Link: GetSchemaType(&propertyValue),
			Schema: GetConceptType(&propertyValue),
		}
		// If the type is a complex type
		fields = append(fields, field)
	}

	return &ConceptDefinition{
		schema: s,
		templateParams: slate.ConceptTemplateParams{
			Name: GetSimpleName(name),
			Fields: fields,
		},
	}
}

func IsComplexType(s *spec.Schema) bool {
	if s == nil {
		return false
	}
	return len(s.SchemaProps.Ref.GetPointer().String()) > 0
}

func GetConceptType(s *spec.Schema) string {
	// Get the reference for complex types
	if IsComplexType(s) {
		s := fmt.Sprintf("%s", s.SchemaProps.Ref.GetPointer())
		s = strings.Replace(s, "/definitions/", "", 1)
		return s
	}
	// Recurse if type is array
	if IsArrayType(s) {
		return GetConceptType(s.Items.Schema)
	}
	return ""
}

func GetParameterType(p *spec.Parameter) string {
	if p.Schema == nil {
		return p.Type
	} else {
		return GetSchemaType(p.Schema)
	}
}

func GetSchemaType(s *spec.Schema) string {
	// Get the reference for complex types
	if IsComplexType(s) {
		s := fmt.Sprintf("%s", s.SchemaProps.Ref.GetPointer())
		s = strings.Replace(s, "/definitions/", "", 1)
		simpleName := GetSimpleName(s)
		s = fmt.Sprintf("[%s](#%s)", simpleName, strings.ToLower(simpleName))
		return s
	}
	// Recurse if type is array
	if IsArrayType(s) {
		return fmt.Sprintf("%s array", GetSchemaType(s.Items.Schema))
	}
	// Get the value for primitive types
	if len(s.Type) > 0 {
		return fmt.Sprintf("%s", s.Type[0])
	}
	panic(fmt.Errorf("No type found for object %v", *s))
}

func IsArrayType(s *spec.Schema) bool {
	if s == nil {
		return false
	}
	return s.Type[0] == "array"
}

func AddOperations(specs []*loads.Document, definitions map[string]*ConceptDefinition, tlc TopLevelConcepts) {
	for _, d := range specs {
		for path, pathItem := range d.Spec().Paths.Paths {
			var concept *ConceptDefinition = nil

			// Get the top level concept this operation is for
			for _, name := range tlc {
				parts := strings.Split(strings.ToLower(name), ".")
				match := fmt.Sprintf("%s/(.*/)?%s(e)?(s)?/", parts[0], parts[1])
				re := regexp.MustCompile(match)
				if re.MatchString(path + "/") {
					if c, f := definitions[name]; !f {
						continue
					} else {
						// Skip this since it matches everything.  Figure this out later.
						if c.templateParams.Name == "Namespace" {
							continue
						}
						concept = c
						break
					}
				}
			}

			// Continue if no top level concepts found
			if concept == nil {
				continue
			}

			pathParams := []spec.Parameter{}
			for _, p := range pathItem.Parameters {
				pathParams = append(pathParams, p)
			}
			if pathItem.Get != nil {
				if o := CreateOperation("GET", path, pathItem.Get, pathParams, &concept.templateParams); o != nil {
					concept.templateParams.Operations = append(concept.templateParams.Operations, *o)
				}
			}
			if pathItem.Post != nil {
				if o := CreateOperation("POST", path, pathItem.Post, pathParams, &concept.templateParams); o != nil {
					concept.templateParams.Operations = append(concept.templateParams.Operations, *o)
				}
			}
			if pathItem.Delete != nil {
				if o := CreateOperation("DELETE", path, pathItem.Delete, pathParams, &concept.templateParams); o != nil {
					concept.templateParams.Operations = append(concept.templateParams.Operations, *o)
				}
			}
			if pathItem.Put != nil {
				if o := CreateOperation("PUT", path, pathItem.Put, pathParams, &concept.templateParams); o != nil {
					concept.templateParams.Operations = append(concept.templateParams.Operations, *o)
				}
			}
			if pathItem.Patch != nil {
				if o := CreateOperation("PATCH", path, pathItem.Patch, pathParams, &concept.templateParams); o != nil {
					concept.templateParams.Operations = append(concept.templateParams.Operations, *o)
				}
			}
		}
	}
	for _, c := range definitions {
		sort.Sort(c.templateParams.Operations)
		pruned := slate.Operations{}
		last := slate.Operation{}
		for _, o := range c.templateParams.Operations {
			if o.Name == last.Name {
				continue
			}
			pruned = append(pruned, o)
			last = o
		}
		c.templateParams.Operations = pruned
	}
}

func CreateOperation(t string, path string, op *spec.Operation, pathParams []spec.Parameter, c *slate.ConceptTemplateParams) *slate.Operation {
	if op == nil || len(op.ID) < 1 {
		return nil
	}
	opQueryParams := []slate.QueryParam{}
	for _, p := range op.Parameters {
		param := slate.QueryParam{
			Description: p.Description,
			Name: p.Name,
			Schema: GetParameterType(&p),
		}
		opQueryParams = append(opQueryParams, param)
	}

	opPathParams := []slate.PathParam{}
	for _, p := range pathParams {
		param := slate.PathParam{
			Description: p.Description,
			Name: p.Name,
			Schema: GetParameterType(&p),
		}
		opPathParams = append(opPathParams, param)
	}

	responses := slate.HttpResponses{}
	for code, response := range op.Responses.ResponsesProps.StatusCodeResponses {
		value := slate.HttpResponse{
			Code: fmt.Sprintf("%d", code),
			Description: response.Description,
			Schema: GetSchemaType(response.Schema),
		}
		responses = append(responses, value)
	}

	return &slate.Operation{
		Description: op.Description,
		Name: op.ID,
		HttpRequest: fmt.Sprintf("%s %s", t, path),
		PathParams: opPathParams,
		QueryParams: opQueryParams,
		HttpResponses: responses,
	}
}

type ConceptDefinition struct {
	schema spec.Schema
	templateParams slate.ConceptTemplateParams
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
