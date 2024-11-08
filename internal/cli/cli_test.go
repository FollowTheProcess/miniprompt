package cli_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/FollowTheProcess/miniprompt/internal/cli"
	"github.com/FollowTheProcess/test"
)

func TestSmoke(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	args := []string{"prompt"}

	cmd, err := cli.Build(context.Background(), stdout, stderr, args)
	test.Ok(t, err) // Command failed to build

	err = cmd.Execute()
	test.Ok(t, err) // Command should not error
}
