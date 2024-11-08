// Package module implements the various modules or segments that go into a shell prompt.
package module

import (
	"os"
	"path/filepath"
	"strings"
)

const symbol = "➜"

// Symbol returns the prompt symbol.
func Symbol() string {
	return symbol
}

// CWD returns the current directory information.
//
// $HOME is truncated to ~/ and if the current directory is
// a git repo the name will be truncated to the repo name.
func CWD(home, cwd string) string {
	info, err := os.Stat(filepath.Join(cwd, ".git"))
	if err == nil && info.IsDir() {
		return filepath.Base(cwd)
	}

	// If $HOME is in the path, replace it otherwise just returns cwd as is
	return "~" + strings.TrimPrefix(cwd, home)
}
