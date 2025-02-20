package repl

import (
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/builtins"
	"github.com/codecrafters-io/shell-starter-go/pkg/utils"
)

const (
	CD   = "cd"
	PWD  = "pwd"
	ECHO = "echo"
	TYPE = "type"
	EXIT = "exit 0"
)

type shellProgram struct {
	repl Repl
}

// NewShellProgram creates a new shell program.
//
// It takes a REPL as a parameter and returns a Program that reads commands
// from the REPL and executes them until the user enters "exit 0". It
// handles shell builtins and external commands differently.
func NewShellProgram(r Repl) Program {
	return &shellProgram{repl: r}
}

// Run starts the shell program's read-eval-print loop.
//
// It reads commands from the REPL, parses them, and executes them. If the
// command is a shell builtin, it is handled internally. Otherwise, it is run as
// an external command. The loop exits when the user enters "exit 0".
func (p *shellProgram) Run() {
outerLoop:
	for {
		input, _ := p.repl.Read()

		switch input {
		case PWD:
			p.execPwdCommand()
			continue
		case EXIT:
			break outerLoop
		}

		args := strings.Split(input, " ")
		if args[0] == CD {
			p.execCDCommand(args[1])
			continue
		}

		_, err := utils.FindFile(args[0])
		if err != nil {
			if len(args) == 2 && args[0] == TYPE {
				p.execTypeCommand(args[1])
				continue
			}
			if args[0] == ECHO {
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

func (s *shellProgram) execTypeCommand(arg string) {
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

func (s *shellProgram) execPwdCommand() {
	output := utils.GetWorkingDirectory()
	s.repl.Print(output, true)
}

func (s *shellProgram) execCDCommand(path string) {
	err := utils.ChangeDirectory(path)
	if err != nil {
		s.repl.Print("cd: "+path+": No such file or directory", true)
	}
}
