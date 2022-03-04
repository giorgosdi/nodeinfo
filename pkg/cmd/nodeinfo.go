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
	"fmt"
	"os"

	"github.com/giorgosdi/nodeinfo/pkg/auth"
	"github.com/giorgosdi/nodeinfo/pkg/options"
	"github.com/giorgosdi/nodeinfo/pkg/resources"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var nodeInfoExample = `
# Get info for a node
kubectl nodeinfo ip-10-200-195-62.example.node


# Get info for a node with a specific KUBECONFIG (flag)
kubectl nodeinfo ip-10-200-195-62.example.node --kubeconfig /path/to/kubeconfig

# Get info for a node with a specific KUBECONFIG (env variable)
export KUBECONFIG=/path/to/kubeconfig;

kubectl nodeinfo ip-10-200-195-62.example.node

# Get info for a node for a specific namespace
kubectl nodeinfo ip-10-200-195-62.example.node -n default
`

func NewNodeInfoCommand(streams genericclioptions.IOStreams) *cobra.Command {

	o := options.NewNodeInfoOptions(streams)

	// cmd represents the nodeinfo command
	var cmd = &cobra.Command{
		Use:          "kubectl nodeinfo <node> [flags]",
		Short:        "Information about a given node",
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		Example:      nodeInfoExample,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			o.Complete(cmd, args)
			_, err := os.Stat(o.Kubeconfig)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			getInfo(cmd, o, args)
		},
	}
	o.Config.AddFlags(cmd.Flags())

	return cmd
}

func getInfo(cmd *cobra.Command, o *options.NodeInfoOptions, args []string) {
	corev1, metricsClient := auth.GetClients(o.Kubeconfig)
	resources.GetPodInfo(o, corev1, metricsClient)
}
