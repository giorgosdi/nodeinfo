package options

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type NodeInfoOptions struct {
	Namespace  string
	Kubeconfig string
	Config     *genericclioptions.ConfigFlags
	Metrics    bool
	Args       []string

	genericclioptions.IOStreams
}

func NewNodeInfoOptions(streams genericclioptions.IOStreams) *NodeInfoOptions {
	return &NodeInfoOptions{
		Config: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

func (o *NodeInfoOptions) Complete(c *cobra.Command, args []string) error {
	var err error
	homeDir, _ := os.UserHomeDir()
	o.Namespace, err = c.Flags().GetString("namespace")
	if err != nil {
		return err
	}
	o.Kubeconfig = fmt.Sprintf("%s/.kube/config", homeDir)

	if value, _ := c.Flags().GetString("kubeconfig"); value != "" {
		o.Kubeconfig, err = c.Flags().GetString("kubeconfig")
		if err != nil {
			return err
		}
	}
	if os.Getenv("KUBECONFIG") != "" {
		o.Kubeconfig = os.Getenv("KUBECONFIG")
	}
	fmt.Println(o.Kubeconfig)

	o.Args = args
	return nil
}
