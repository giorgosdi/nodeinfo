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
	"fmt"
	"os"

	"nodeinfo/pkg/auth"
	"nodeinfo/pkg/resources"

	"github.com/spf13/cobra"

	flag "github.com/spf13/pflag"
)

type Flags struct {
	namespace string
	node      string
	config    string
}

// nodeinfoCmd represents the nodeinfo command
var nodeInfoCmd = &cobra.Command{
	Use:     "nodeinfo",
	Aliases: []string{"nf", "info"},
	Short:   "Info about a given node",
	Run: func(cmd *cobra.Command, args []string) {
		getInfo()
	},
}

func init() {
	homeDir, _ := os.UserHomeDir()
	defaultKubeConfig := fmt.Sprintf("%s/.kube/config", homeDir)
	fmt.Println(defaultKubeConfig)
	flag.StringP("namespace", "n", "", "The namespace you want search in")
	flag.String("node", "", "The node you want to query")
	flag.String("config", defaultKubeConfig, "KUBECONFIG you want to use")
	flag.Parse()
}

func flags() Flags {
	ns := flag.Lookup("namespace").Value.String()
	node := flag.Lookup("node").Value.String()
	conf := flag.Lookup("config").Value.String()

	return Flags{
		namespace: ns,
		node:      node,
		config:    conf,
	}
}

func getInfo() {

	flags := flags()

	client := auth.GetConfig(flags.config)
	resources.GetPodInfo(flags.namespace, flags.node, client)
}
