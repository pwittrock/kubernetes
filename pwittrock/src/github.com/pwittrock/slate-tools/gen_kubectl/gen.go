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

package gen_kubectl

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/pwittrock/slate-tools/lib"
	"gopkg.in/yaml.v2"
)

func GenerateSlateFiles() {
	spec := KubectlSpec{}

	if len(*lib.YamlFile) < 1 {
		fmt.Printf("Must specify --yaml-file.\n")
		os.Exit(2)
	}

	contents, err := ioutil.ReadFile(*lib.YamlFile)
	if err != nil {
		fmt.Printf("Failed to read yaml file %s: %v", *lib.YamlFile, err)
	}

	err = yaml.Unmarshal(contents, &spec)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	NormalizeSpec(&spec)

	WriteIndexFile(spec)
	WriteCommandFiles(spec)
}

func NormalizeSpec(spec *KubectlSpec) {
	for _, g  := range spec.TopLevelCommandGroups {
		for _, c := range g.Commands {
			FormatCommand(c.Command)
			for _, s := range c.SubCommands {
				FormatCommand(s)
			}
		}
	}
}

func FormatCommand(c *Command) {
	c.Example = FormatExample(c.Example)
	c.Description = FormatDescription(c.Description)
}

func FormatDescription(input string) string {
	return strings.Replace(input, "   *", "*", 10000)
}

func FormatExample(input string) string {
	last := ""
	result := ""
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) < 1 {
			continue
		}
		if strings.HasPrefix(line, "#") {
			if len(strings.TrimSpace(strings.Replace(line, "#", ">", 1))) < 1 {
				continue
			}
		}

		if strings.HasPrefix(line, "#") {
			if last == "command" {
				result += "\n```\n\n"
			}
			if last == "comment" {
				result += " " + line
			} else {
				result += strings.Replace(line, "#", ">", 1)
			}
			last = "comment"
		} else {
			if last != "command" {
				result += "\n\n```shell"
			}
			result += "\n" + line
			last = "command"
		}
	}

	// Close the final command if needed
	if last == "command" {
		result += "\n```\n"
	}
	return result
}

func WriteIndexFile(spec KubectlSpec) {
	f, err := os.Create(*lib.BuildDir + "/index.html.md")
	if err != nil {
		fmt.Printf("Failed to open index: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	t, err := template.New("index.template").ParseFiles(*lib.TemplateDir + "/index.template")
	if err != nil {
		fmt.Printf("Failed to parse template: %v", err)
		os.Exit(1)
	}

	// Write the instantiated template to the index file
	err = t.Execute(f, spec)
	if err != nil {
		fmt.Printf("Failed to execute template: %v", err)
		os.Exit(1)
	}
}

func WriteCommandFiles(params KubectlSpec) {
	t, err := template.New("command.template").ParseFiles(*lib.TemplateDir + "/command.template")
	if err != nil {
		fmt.Printf("Failed to parse template: %v", err)
		os.Exit(1)
	}

	for _, g := range params.TopLevelCommandGroups {
		for _, tlc := range g.Commands {
			WriteCommandFile(t, tlc)
		}
	}
}

func WriteCommandFile(t *template.Template, params TopLevelCommand) {
	f, err := os.Create(*lib.BuildDir + "/includes/_generated_" + strings.ToLower(params.Command.Name) + ".md")
	if err != nil {
		fmt.Printf("Failed to open index: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	err = t.Execute(f, params)
	if err != nil {
		fmt.Printf("Failed to execute template: %v", err)
		os.Exit(1)
	}
}
