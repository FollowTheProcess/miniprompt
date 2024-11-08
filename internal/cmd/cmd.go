// Package cmd implements miniprompt's command line interface.
package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/FollowTheProcess/cli"
	"github.com/FollowTheProcess/miniprompt/internal/module"
)

var (
	version   = "dev" // miniprompt's version number, set at compile time
	commit    = ""    // The SHA1 of the commit miniprompt was built from, set at compile time
	buildDate = ""    // The build timestamp, set at compile time
)

// ExecutionContext is the context in which miniprompt is being run, it includes
// things like input/output streams, CWD etc.
type ExecutionContext struct {
	CWD    string    // The current working directory
	Home   string    // The user's home directory
	Stdout io.Writer // Writer for the prompt
	Stderr io.Writer // Logging or error output
	Args   []string  // Command line args
}

// Context returns the ExecutionContext miniprompt is being run in, taking
// values from the operating system.
func Context() (ExecutionContext, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return ExecutionContext{}, fmt.Errorf("could not get $HOME: %w", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return ExecutionContext{}, fmt.Errorf("could not get $CWD: %w", err)
	}

	return ExecutionContext{
		CWD:    cwd,
		Home:   home,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Args:   os.Args[1:], // Trim the program off (first arg)
	}, nil
}

// Build returns the miniprompt CLI.
func Build(ctx context.Context, ex ExecutionContext) (*cli.Command, error) {
	prompt := func() (*cli.Command, error) {
		return cli.New(
			"prompt",
			cli.Short("Show the prompt"),
			cli.Run(func(cmd *cli.Command, args []string) error {
				fmt.Fprintf(cmd.Stdout(), "\n%s\n%s", module.CWD(ex.Home, ex.CWD), module.Symbol())
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
		cli.Stdout(ex.Stdout),
		cli.Stderr(ex.Stderr),
		cli.SubCommands(prompt),
		cli.OverrideArgs(ex.Args),
	)
}
