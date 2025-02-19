package repl

import (
	"os/exec"
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

func (p *shellProgram) Run() {
	for {
		input, _ := p.repl.Read()
		if input == "exit 0" {
			break
		}
		args := strings.Split(input, " ")
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
}

func (s *shellProgram) echo(value []string) {
	s.repl.Print(strings.Join(value, " "), true)
}

func (s *shellProgram) executeTypeCommand(arg string) {
	var output []string
	var outputString string

	exists := builtins.IsBuiltin(arg)
	if !exists {
		path, err := findFile(arg)
		if err != nil {
			s.repl.Print(arg, false)
			return
		}
		output = []string{arg, "is", path}
		outputString = buildStrings(output)
		s.repl.Print(outputString, true)
		return
	}

	output = []string{arg, "is a shell builtin"}
	outputString = buildStrings(output)
	s.repl.Print(outputString, true)
}

func findFile(filename string) (string, error) {
	path, err := exec.LookPath(filename)
	if err != nil {
		return "", err
	}
	return path, nil
}

func buildStrings(args []string) string {
	var response strings.Builder
	for _, arg := range args {
		response.WriteString(arg)
		response.WriteString(" ")
	}
	return response.String()
}
