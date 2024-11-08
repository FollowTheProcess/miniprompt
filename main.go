package main

import (
	"context"
	"fmt"
	"os"

	"github.com/FollowTheProcess/miniprompt/internal/cli"
	"github.com/FollowTheProcess/msg"
)

func main() {
	if err := run(); err != nil {
		msg.Error("%s", err)
		os.Exit(1)
	}
}

func run() error {
	cmd, err := cli.Build(context.Background(), os.Stdout, os.Stderr, os.Args)
	if err != nil {
		return fmt.Errorf("could not build miniprompt CLI: %w", err)
	}
	return cmd.Execute()
}
