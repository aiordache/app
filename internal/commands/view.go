package commands

import (
	"fmt"

	"github.com/docker/app/internal"
	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

type viewOptions struct {
	parametersOptions
	registryOptions
	pullOptions
}

func viewCmd(dockerCli command.Cli) *cobra.Command {
	var opts viewOptions
	cmd := &cobra.Command{
		Use:     "view [APP_NAME] [OPTIONS]",
		Short:   "Shows metadata, parameters and a summary of the Compose file for a given application",
		Example: `$ docker app view myapp.dockerapp`,
		Args:    cli.RequiresMaxArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runView(dockerCli, firstOrEmpty(args), opts)
		},
	}
	opts.parametersOptions.addFlags(cmd.Flags())
	opts.registryOptions.addFlags(cmd.Flags())
	opts.pullOptions.addFlags(cmd.Flags())
	return cmd
}

func runView(dockerCli command.Cli, appname string, opts viewOptions) error {
	defer muteDockerCli(dockerCli)()
	action, installation, errBuf, err := prepareCustomAction(internal.ActionViewName, dockerCli, appname, nil, opts.registryOptions, opts.pullOptions, opts.parametersOptions)
	if err != nil {
		return err
	}
	if err := action.Run(&installation.Claim, nil, nil); err != nil {
		return fmt.Errorf("view failed: %s\n%s", err, errBuf)
	}
	return nil
}
