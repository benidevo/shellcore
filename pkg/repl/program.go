package repl

import (
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/builtins"
	"github.com/codecrafters-io/shell-starter-go/pkg/utils"
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

// Run starts the shell program.
//
// It reads commands from the REPL and executes them until the user enters
// "exit 0". It handles shell builtins and external commands differently.
func (p *shellProgram) Run() {
	for {
		input, _ := p.repl.Read()
		if input == "exit 0" {
			break
		}

		args := strings.Split(input, " ")
		_, err := utils.FindFile(args[0])
		if err != nil {
			if len(args) == 2 && args[0] == "type" {
				p.executeTypeCommand(args[1])
				continue
			}
			if args[0] == "echo" {
				p.echo(args[1:])

			} else {
				p.repl.Print(input, false)
			}
		}

		utils.ExecuteScript(args[0], args[1:]...)

	}
}

func (s *shellProgram) echo(value []string) {
	s.repl.Print(strings.Join(value, " "), true)
}

func (s *shellProgram) executeTypeCommand(arg string) {
	var output []string
	var outputString string

	exists := builtins.IsBuiltin(arg)
	if !exists {
		path, err := utils.FindFile(arg)
		if err != nil {
			s.repl.Print(arg, false)
			return
		}
		output = []string{arg, "is", path}
		outputString = utils.BuildStrings(output)
		s.repl.Print(outputString, true)
		return
	}

	output = []string{arg, "is a shell builtin"}
	outputString = utils.BuildStrings(output)
	s.repl.Print(outputString, true)
}
