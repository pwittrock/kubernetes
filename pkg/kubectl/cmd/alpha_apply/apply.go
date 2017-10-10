/*
Copyright 2014 The Kubernetes Authors.

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

package alpha_apply

import (
	"fmt"
	"io"
	"time"

	"github.com/spf13/cobra"

	"encoding/json"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/kubectl/apply/parse"
	"k8s.io/kubernetes/pkg/kubectl/apply/strategy"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/resource"
	"k8s.io/kubernetes/pkg/kubectl/util/i18n"
	"os"
	"path/filepath"
)

type ApplyOptions struct {
	FilenameOptions resource.FilenameOptions
	Selector        string
	Force           bool
	Prune           bool
	Cascade         bool
	GracePeriod     int
	Timeout         time.Duration
	cmdBaseName     string
}

var (
	applyLong    = templates.LongDesc(i18n.T(``))
	applyExample = templates.Examples(i18n.T(``))
)

func NewCmdApply(baseName string, f cmdutil.Factory, out, errOut io.Writer) *cobra.Command {
	var options ApplyOptions

	// Store baseName for use in printing warnings / messages involving the base command name.
	// This is useful for downstream command that wrap this one.
	options.cmdBaseName = baseName

	cmd := &cobra.Command{
		Use:     "apply -f FILENAME",
		Short:   i18n.T("Apply a configuration to a resource by filename or stdin"),
		Long:    applyLong,
		Example: applyExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(validateArgs(cmd, args))
			cmdutil.CheckErr(RunApply(f, cmd, out, errOut, &options))
		},
	}

	usage := "that contains the configuration to apply"
	cmdutil.AddFilenameOptionFlags(cmd, &options.FilenameOptions, usage)
	cmd.MarkFlagRequired("filename")
	cmd.Flags().Bool("overwrite", true, "Automatically resolve conflicts between the modified and live configuration by using values from the modified configuration")
	cmd.Flags().BoolVar(&options.Prune, "prune", false, "Automatically delete resource objects, including the uninitialized ones, that do not appear in the configs and are created by either apply or create --save-config. Should be used with either -l or --all.")
	cmd.Flags().BoolVar(&options.Cascade, "cascade", true, "Only relevant during a prune or a force apply. If true, cascade the deletion of the resources managed by pruned or deleted resources (e.g. Pods created by a ReplicationController).")
	cmd.Flags().IntVar(&options.GracePeriod, "grace-period", -1, "Only relevant during a prune or a force apply. Period of time in seconds given to pruned or deleted resources to terminate gracefully. Ignored if negative.")
	cmd.Flags().DurationVar(&options.Timeout, "timeout", 0, "Only relevant during a force apply. The length of time to wait before giving up on a delete of the old resource, zero means determine a timeout from the size of the object. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h).")
	cmdutil.AddValidateFlags(cmd)
	cmd.Flags().StringVarP(&options.Selector, "selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	cmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types.")
	cmd.Flags().StringArray("prune-whitelist", []string{}, "Overwrite the default whitelist with <group/version/kind> for --prune")
	cmdutil.AddDryRunFlag(cmd)
	cmdutil.AddPrinterFlags(cmd)
	cmdutil.AddRecordFlag(cmd)
	cmdutil.AddInclude3rdPartyFlags(cmd)
	cmdutil.AddIncludeUninitializedFlag(cmd)

	return cmd
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return cmdutil.UsageErrorf(cmd, "Unexpected args: %v", args)
	}

	return nil
}

func RunApply(f cmdutil.Factory, cmd *cobra.Command, out, errOut io.Writer, options *ApplyOptions) error {
	schema, err := f.Validator(cmdutil.GetFlagBool(cmd, "validate"))
	if err != nil {
		return err
	}

	cmdNamespace, enforceNamespace, err := f.DefaultNamespace()
	if err != nil {
		return err
	}

	mapper, typer, err := f.UnstructuredObject()
	if err != nil {
		return err
	}

	// include the uninitialized objects by default if --prune is true
	// unless explicitly set --include-uninitialized=false
	includeUninitialized := cmdutil.ShouldIncludeUninitialized(cmd, options.Prune)
	r := f.NewBuilder().
		Unstructured(f.UnstructuredClientForMapping, mapper, typer).
		Schema(schema).
		ContinueOnError().
		NamespaceParam(cmdNamespace).DefaultNamespace().
		FilenameParam(enforceNamespace, &options.FilenameOptions).
		SelectorParam(options.Selector).
		IncludeUninitialized(includeUninitialized).
		Flatten().
		Do()
	err = r.Err()
	if err != nil {
		return err
	}

	encoder := f.JSONEncoder()
	decoder := f.Decoder(false)

	err = r.Visit(func(info *resource.Info, err error) error {
		if err != nil {
			return err
		}

		dir, err := ioutil.TempDir("", "kubectl-apply-diff")
		if err != nil {
			return err
		}

		if err := info.Get(); err != nil {
			if !errors.IsNotFound(err) {
				return cmdutil.AddSourceToErr(fmt.Sprintf("retrieving current configuration of:\n%v\nfrom server for:", info), info.Source, err)
			}
			return err
		}
		return write(f, cmd, encoder, decoder, dir, info)
	})

	if err != nil {
		return err
	}
	return nil
}

func write(f cmdutil.Factory, cmd *cobra.Command, e runtime.Encoder, d runtime.Decoder, dir string, info *resource.Info) error {
	mapping := info.ResourceMapping()
	printer, err := f.PrinterForMapping(cmd, false, nil, mapping, false)
	if err != nil {
		return err
	}

	// Get openapi
	resources, err := f.OpenAPISchema()
	if err != nil {
		return err
	}
	elementParser := parse.Factory{resources}

	// Write local file for diff
	localBytes, err := runtime.Encode(f.JSONEncoder(), info.VersionedObject)
	if err != nil {
		return err
	}
	local := map[string]interface{}{}
	err = json.Unmarshal(localBytes, &local)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filepath.Join(dir, "local"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return err
	}
	err = printer.PrintObj(info.VersionedObject, file)
	if err != nil {
		return err
	}
	file.Close()
	fmt.Printf("local file: %s\n", filepath.Join(dir, "local"))

	// Write remote file for diff
	remoteBytes, err := runtime.Encode(f.JSONEncoder(), info.Object)
	if err != nil {
		return err
	}
	remote := map[string]interface{}{}
	err = json.Unmarshal(remoteBytes, &remote)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(filepath.Join(dir, "remote"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return err
	}
	err = printer.PrintObj(info.Object, file)
	if err != nil {
		return err
	}
	file.Close()
	fmt.Printf("remote file: %s\n", filepath.Join(dir, "remote"))

	// Write last applied file for diff
	accessor, err := meta.Accessor(info.Object)
	if err != nil {
		return err
	}
	annots := accessor.GetAnnotations()
	if annots == nil {
		annots = map[string]string{}
	}
	lastBytes := annots[api.LastAppliedConfigAnnotation]
	last := map[string]interface{}{}
	err = json.Unmarshal([]byte(lastBytes), &last)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(filepath.Join(dir, "last"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	lastObj, _, err := d.Decode([]byte(lastBytes), nil, nil)
	if err != nil {
		return err
	}
	err = printer.PrintObj(lastObj, file)
	if err != nil {
		return err
	}
	file.Close()
	fmt.Printf("last file: %s\n", filepath.Join(dir, "last"))

	// Write merged file for diff
	obj, err := elementParser.CreateElement(last, local, remote)
	if err != nil {
		return err
	}

	result, err := obj.Merge(strategy.Create(strategy.Options{}))
	merged := result.MergedResult
	mergedBytes, err := json.Marshal(merged)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(filepath.Join(dir, "merged"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	lastObj, _, err = d.Decode([]byte(mergedBytes), nil, nil)
	if err != nil {
		return err
	}
	err = printer.PrintObj(lastObj, file)
	if err != nil {
		return err
	}
	file.Close()
	fmt.Printf("merged file: %s\n", filepath.Join(dir, "merged"))

	return nil
}
