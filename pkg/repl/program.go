package repl

import "strings"

type ShellProgram struct {
	repl Repl
}

func NewShellProgram(r Repl) *ShellProgram {
	return &ShellProgram{repl: r}
}

func (p *ShellProgram) Run() {
	for {
		input, _ := p.repl.Read()
		if input == "exit 0" {
			break
		}
		args := strings.Split(input, " ")
		if args[0] == "echo" {
			p.echo(args[1:])

		} else {
			p.repl.Print(input)
		}
	}
}

func (s *ShellProgram) echo(value []string) {
	s.repl.Print(strings.Join(value, " "))
}
