package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ShellRepl struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

// NewShellRepl creates a new REPL for the shell.
//
// It creates a *bufio.Reader for reading from os.Stdin and a *bufio.Writer for
// writing to os.Stdout.
func NewShellRepl() *ShellRepl {
	return &ShellRepl{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

// Read displays the shell prompt and waits for user input.
//
// It reads a line from the standard input and returns the input string
// without the trailing newline character. If an error occurs during
// reading, it returns an empty string and the encountered error.
func (s *ShellRepl) Read() (string, error) {
	s.writeShellPrompt()

	input, err := s.reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		return "", err
	}

	return strings.TrimSuffix(input, "\n"), nil
}

// Print writes the given output string to the shell's writer.
//
// If the output is not valid (i.e. the command is not recognized), it writes
// "<command>: not found" followed by a newline. Otherwise, it writes just the
// output string followed by a newline.
func (s *ShellRepl) Print(output string, valid bool) {
	if !valid {
		fmt.Fprintf(s.writer, "%s: not found\n", output)
	} else {
		fmt.Fprintf(s.writer, "%s\n", output)
	}
	s.writer.Flush()
}

func (s *ShellRepl) writeShellPrompt() {
	s.writer.WriteString("$ ")
	s.writer.Flush()
}
