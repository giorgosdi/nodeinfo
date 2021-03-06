/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"os"

	"github.com/giorgosdi/nodeinfo/pkg/cmd"

	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func KubectlNodeinfo() {
	flags := pflag.NewFlagSet("kubectl-nodeinfo", pflag.ExitOnError)
	pflag.CommandLine = flags
	nodeInfoCmd := cmd.NewNodeInfoCommand(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := nodeInfoCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
