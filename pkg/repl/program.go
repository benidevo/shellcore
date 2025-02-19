package repl

import (
	"os/exec"
	"strings"
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
	path, err := findFile(arg)
	if err != nil {
		s.repl.Print(arg, false)
		return
	}
	response := strings.Builder{}
	response.WriteString(arg)
	response.WriteString(" is ")
	response.WriteString(path)
	s.repl.Print(response.String(), true)
}

func findFile(filename string) (string, error) {
	path, err := exec.LookPath(filename)
	if err != nil {
		return "", err
	}
	return path, nil
}
