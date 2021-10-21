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
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	flag "github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// nodeinfoCmd represents the nodeinfo command
var nodeInfoCmd = &cobra.Command{
	Use:     "nodeinfo",
	Aliases: []string{"nf", "info"},
	Short:   "Info about a given node",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nodeinfo called")
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

func getInfo() {
	ns := flag.Lookup("namespace").Value.String()
	node := flag.Lookup("node").Value.String()
	conf := flag.Lookup("config").Value.String()

	config, _ := clientcmd.BuildConfigFromFlags("", conf)
	clientset, _ := kubernetes.NewForConfig(config)
	pods, _ := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	fmt.Printf("Node : %s\n", node)
	for _, pod := range pods.Items {
		if pod.Spec.NodeName == node {
			fmt.Printf("POD: %s\n", pod.ObjectMeta.Name)
			for _, container := range pod.Spec.Containers {
				fmt.Printf("Container: %s\n", container.Name)
				fmt.Printf("Requested: %s   |    Limit: %s\n", container.Resources.Requests.Cpu(), container.Resources.Limits.Cpu())
			}
		}
	}
}
