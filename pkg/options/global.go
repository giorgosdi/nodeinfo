package options

import (
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
	o.Namespace, err = c.Flags().GetString("namespace")
	if err != nil {
		return err
	}
	o.Kubeconfig, err = c.Flags().GetString("kubeconfig")
	if err != nil {
		return err
	}
	o.Args = args
	return nil
}
