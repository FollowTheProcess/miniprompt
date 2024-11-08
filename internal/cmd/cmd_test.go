package cmd_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/FollowTheProcess/miniprompt/internal/cmd"
	"github.com/FollowTheProcess/miniprompt/internal/module"
	"github.com/FollowTheProcess/test"
)

func TestSmoke(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	testContext := cmd.ExecutionContext{
		CWD:    "/somewhere/user/project",
		Home:   "/somewhere/user",
		Stdout: stdout,
		Stderr: stderr,
		Args:   []string{"prompt"},
	}

	cmd, err := cmd.Build(context.Background(), testContext)
	test.Ok(t, err) // Command failed to build

	err = cmd.Execute()
	test.Ok(t, err) // Command should not error

	test.Equal(t, stdout.String(), fmt.Sprintf("\n~/project\n%s", module.Symbol()))
}
