// Package cmd implements miniprompt's command line interface.
package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/FollowTheProcess/cli"
)

var (
	version   = "dev" // miniprompt's version number, set at compile time
	commit    = ""    // The SHA1 of the commit miniprompt was built from, set at compile time
	buildDate = ""    // The build timestamp, set at compile time
)

// Build returns the miniprompt CLI.
func Build(ctx context.Context, stdout, stderr io.Writer, args []string) (*cli.Command, error) {
	prompt := func() (*cli.Command, error) {
		return cli.New(
			"prompt",
			cli.Short("Show the prompt"),
			cli.Run(func(cmd *cli.Command, args []string) error {
				fmt.Fprintln(cmd.Stdout(), "prompt showy showy")
				return nil
			}),
		)
	}
	return cli.New(
		"miniprompt",
		cli.Short("A teeny tiny shell prompt"),
		cli.Version(version),
		cli.Commit(commit),
		cli.BuildDate(buildDate),
		cli.Allow(cli.NoArgs()),
		cli.Stdout(stdout),
		cli.Stderr(stderr),
		cli.SubCommands(prompt),
		cli.OverrideArgs(args),
	)
}
