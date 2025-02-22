package utils

import (
	"os"
	"os/exec"
)

// ExecuteScript runs the given script with the given arguments.
//
// It sets the script's stdout and stderr to os.Stdout and os.Stderr
// respectively, and returns the error from running the command.
func ExecuteScript(script string, args ...string) error {
	cmd := exec.Command(script, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
