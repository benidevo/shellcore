package repl

import (
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/builtins"
)

type shellProgram struct {
	repl Repl
}

// NewShellProgram creates a new shell program for the given REPL.
//
// It returns a *shellProgram which implements the Program interface.
func NewShellProgram(r Repl) *shellProgram {
	return &shellProgram{repl: r}
}

// Run starts the shell program's main loop, reading user input and executing commands.
//
// It reads input from the REPL in a loop, exiting when "exit 0" is entered.
// If the command is "type" followed by a single argument, it checks if the argument
// is a shell builtin and prints the result. If the command is "echo", it calls the
// echo method with the provided arguments. For any other command, it prints "command
// not found" if the command is not recognized.
func (p *shellProgram) Run() {
	for {
		input, _ := p.repl.Read()
		if input == "exit 0" {
			break
		}
		args := strings.Split(input, " ")
		if len(args) == 2 && args[0] == "type" {
			p.executeTypeCommand(args[1:])
			continue
		}
		if args[0] == "echo" {
			p.echo(args[1:])

		} else {
			p.repl.Print(input, false)
		}
	}
}

func (s *shellProgram) echo(value []string) {
	s.repl.Print(strings.Join(value, " "), true)
}

func (s *shellProgram) executeTypeCommand(args []string) {
	exists := builtins.IsBuiltin(args[0])
	if !exists {
		s.repl.Print(args[0], false)
		return
	}
	response := strings.Builder{}
	response.WriteString(args[0])
	response.WriteString(" is a shell builtin")
	s.repl.Print(response.String(), true)
}
