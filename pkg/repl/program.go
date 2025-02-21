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
	CAT  = "cat"
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
		if input == EXIT {
			break outerLoop
		}
		parsedInput := utils.ParseCommand(input)
		if len(parsedInput) == 0 {
			continue
		}

		command := parsedInput[0]
		switch command {
		case PWD:
			p.execPwdCommand()
			continue
		case CD:
			p.execCdCommand(parsedInput[1])
			continue
		case ECHO:
			p.echo(parsedInput[1:])
			continue
		case CAT:
			parsedInput := utils.ParseCommand(input)
			utils.ExecuteScript(parsedInput[0], parsedInput[1:]...)
			continue
		}

		_, err := utils.FindFile(command)
		if err != nil {
			if len(parsedInput) == 2 && command == TYPE {
				p.execTypeCommand(parsedInput[1])
				continue
			}

			p.repl.Print(input, false)
			continue
		}

		utils.ExecuteScript(command, parsedInput[1:]...)
		continue
	}
}

func (s *shellProgram) echo(value []string) {
	output := strings.Join(value, " ")
	s.repl.Print(output, true)
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

func (s *shellProgram) execCdCommand(path string) {
	err := utils.ChangeDirectory(path)
	if err != nil {
		s.repl.Print("cd: "+path+": No such file or directory", true)
	}
}
