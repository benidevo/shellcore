package utils

import (
	"os"
	"os/exec"
	"strings"
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

// ParseCommand splits the input string into a slice of arguments, handling quoted substrings.
//
// It iterates through each character of the input string, using a strings.Builder to accumulate characters
// for each argument. The function respects single quotes, treating the characters within quotes as a
// single argument, even if spaces are present. It removes surrounding quotes from arguments after parsing.
//
// Returns a slice of strings representing the parsed arguments.
func ParseCommand(input string) []string {
	var args []string
	var currentArg strings.Builder
	inQuotes := false

	for i := 0; i < len(input); i++ {
		char := input[i]
		switch char {
		case '\'':
			inQuotes = !inQuotes
		case ' ':
			if !inQuotes {
				if currentArg.Len() > 0 {
					args = append(args, currentArg.String())
					currentArg.Reset()
				}
			} else {
				currentArg.WriteByte(char)
			}
		default:
			currentArg.WriteByte(char)
		}
	}

	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	for i, arg := range args {
		if len(arg) > 0 && arg[0] == '\'' && arg[len(arg)-1] == '\'' {
			args[i] = arg[1 : len(arg)-1]
		}
	}

	return args
}
