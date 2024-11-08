package main

import (
	"fmt"
	"os"

	"github.com/FollowTheProcess/cli"
	"github.com/FollowTheProcess/msg"
)

var (
	version   = "dev" // miniprompt's version number, set at compile time
	commit    = ""    // The SHA1 of the commit miniprompt was built from, set at compile time
	buildDate = ""    // The build timestamp, set at compile time
)

func main() {
	if err := run(); err != nil {
		msg.Error("%s", err)
		os.Exit(1)
	}
}

func run() error {
	prompt := func() (*cli.Command, error) {
		return cli.New(
			"prompt",
			cli.Short("Show the prompt"),
			cli.Run(func(cmd *cli.Command, args []string) error {
				fmt.Println("show the prompt here")
				return nil
			}),
		)
	}

	cmd, err := cli.New(
		"miniprompt",
		cli.Short("A teeny tiny shell prompt"),
		cli.Version(version),
		cli.Commit(commit),
		cli.BuildDate(buildDate),
		cli.Allow(cli.NoArgs()),
		cli.SubCommands(prompt),
	)
	if err != nil {
		return fmt.Errorf("could not build miniprompt CLI: %w", err)
	}

	return cmd.Execute()
}
