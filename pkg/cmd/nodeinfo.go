/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
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
	"github.com/giorgosdi/nodeinfo/pkg/auth"
	"github.com/giorgosdi/nodeinfo/pkg/options"
	"github.com/giorgosdi/nodeinfo/pkg/resources"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	flag "github.com/spf13/pflag"
)

func NewNodeInfoCommand(streams genericclioptions.IOStreams) *cobra.Command {

	o := options.NewNodeInfoOptions(streams)

	// cmd represents the nodeinfo command
	var cmd = &cobra.Command{
		Use:     "nodeinfo <node>",
		Aliases: []string{"nf", "info"},
		Short:   "Info about a given node",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			o.Complete(cmd, args)
			getInfo(cmd, o, args)
		},
	}

	flag.BoolP("metrics", "m", false, "show metrics")
	o.Config.AddFlags(cmd.Flags())

	return cmd
}

func getInfo(cmd *cobra.Command, o *options.NodeInfoOptions, args []string) {
	corev1, metricsClient := auth.GetClients(o.Kubeconfig)
	resources.GetPodInfo(o, corev1, metricsClient)
}
