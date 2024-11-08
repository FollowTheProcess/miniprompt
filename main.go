package main

import (
	"context"
	"fmt"
	"os"

	"github.com/FollowTheProcess/miniprompt/internal/cmd"
	"github.com/FollowTheProcess/msg"
)

func main() {
	if err := run(); err != nil {
		msg.Error("%s", err)
		os.Exit(1)
	}
}

func run() error {
	executionContext, err := cmd.Context()
	if err != nil {
		return err
	}
	cmd, err := cmd.Build(context.Background(), executionContext)
	if err != nil {
		return fmt.Errorf("could not build miniprompt CLI: %w", err)
	}
	return cmd.Execute()
}
