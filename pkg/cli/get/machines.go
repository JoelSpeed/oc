package get

import (
	"fmt"
	"io"
	"strings"

	"k8s.io/kubectl/pkg/cmd/get"

	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
)

type MachinesOptions struct {
	errOut io.Writer
	// Embed kubectl's GetOptions directly.
	*get.GetOptions
}

func NewMachinesOptions(parent string, streams genericclioptions.IOStreams) *MachinesOptions {
	return &MachinesOptions{
		errOut:     streams.ErrOut,
		GetOptions: get.NewGetOptions(parent, streams),
	}
}

func NewCmdMachines(parent string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewMachinesOptions(parent, streams)

	cmd := get.NewCmdGet(parent, f, streams)
	cmd.Use = fmt.Sprintf("machines [(-o|--output=)%s] (TYPE[.VERSION][.GROUP] [NAME | -l label] | TYPE[.VERSION][.GROUP]/NAME ...) [flags]", strings.Join(o.PrintFlags.AllowedFormats(), "|"))
	cmd.Run = func(cmd *cobra.Command, args []string) {
		args = append([]string{"machines"}, args...)

		kcmdutil.CheckErr(o.Complete(f, cmd, args))
		kcmdutil.CheckErr(o.Validate(cmd))
		kcmdutil.CheckErr(o.Run(f, cmd, args))
	}

	return cmd
}

func (m *MachinesOptions) Run(f kcmdutil.Factory, cmd *cobra.Command, args []string) error {
	fmt.Fprintf(m.errOut, "WARNING: machines is ambiguous, try machines.machine.openshift.io or machines.x-cluster.k8s.io\n")

	return m.GetOptions.Run(f, cmd, args)
}
