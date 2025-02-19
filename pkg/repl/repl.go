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

func NewShellRepl() *ShellRepl {
	return &ShellRepl{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (s *ShellRepl) Read() (string, error) {
	s.writeShellPrompt()

	input, err := s.reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		return "", err
	}

	return strings.TrimSuffix(input, "\n"), nil
}

func (s *ShellRepl) Print(output string) {
	fmt.Fprintf(s.writer, "%s\n", output)
	s.writer.Flush()
}

func (s *ShellRepl) writeShellPrompt() {
	s.writer.WriteString("$ ")
	s.writer.Flush()
}
