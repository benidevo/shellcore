package repl

import (
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/builtins"
)

type ShellProgram struct {
	repl Repl
}

func NewShellProgram(r Repl) *ShellProgram {
	return &ShellProgram{repl: r}
}

// Run starts the shell program and runs the REPL loop.
//
// It reads input from the user, checks if it's an exit command, and
// then splits the input into words to check if it's a valid command.
// If it is, it runs the command, otherwise it prints an error message.
func (p *ShellProgram) Run() {
	for {
		input, _ := p.repl.Read()
		if input == "exit 0" {
			break
		}
		args := strings.Split(input, " ")
		if len(args) == 2 && args[1] == "echo" {
			exists := builtins.IsBuiltin(args[0])
			if !exists {
				p.repl.Print(input, false)
				continue
			}
			response := strings.Builder{}
			response.WriteString(args[0])
			response.WriteString(" is a shell builtin")
			p.repl.Print(response.String(), true)
			continue
		}
		if args[0] == "echo" {
			p.echo(args[1:])

		} else {
			p.repl.Print(input, false)
		}
	}
}

func (s *ShellProgram) echo(value []string) {
	s.repl.Print(strings.Join(value, " "), true)
}
