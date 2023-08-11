package cli

import (
	"context"

	"github.com/urfave/cli/v2"
)

type CLI struct {
	App *cli.App
}

var deployCommand *cli.Command = &cli.Command{
    Name: "deploy",
    Usage: "deploy application using yaml",
    Flags: []cli.Flag{
        &cli.StringFlag{
            Name: "config",
            Aliases: []string{"c"},
            Usage: "Load configuration from `FILE`",
        },
    },
    Action: deployAction,
}

func NewCLI() *CLI {
	return &CLI{
		App: &cli.App{
			Name:  "caps",
			Usage: "Deploy to Azure Container Apps",
            Commands: cli.Commands{
                deployCommand,
            },
		},
	}
}

func (c *CLI) Run(ctx context.Context, args []string) error {
	return c.App.RunContext(ctx, args)
}
