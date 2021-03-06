package boot

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/spf13/cobra"
)

// StepBootOptions contains the command line flags
type StepBootOptions struct {
	opts.StepOptions
}

// NewCmdStepBoot Steps a command object for the "step" command
func NewCmdStepBoot(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &StepBootOptions{
		StepOptions: opts.StepOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "boot",
		Short:   "boot [command]",
		Aliases: []string{},
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	cmd.AddCommand(NewCmdStepBootUpgrade(commonOpts))
	cmd.AddCommand(NewCmdStepBootVault(commonOpts))
	return cmd
}

// Run implements this command
func (o *StepBootOptions) Run() error {
	return o.Cmd.Help()
}
