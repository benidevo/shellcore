package repl

import (
	"strings"

	"github.com/benidevo/shellcore/internal/builtins"
	"github.com/benidevo/shellcore/internal/parser"
	"github.com/benidevo/shellcore/pkg/utils"
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
		if strings.TrimSpace(input) == EXIT {
			break outerLoop
		}
		parsedInput := parser.ParseCommand(input)
		if len(parsedInput) == 0 {
			continue
		}

		command := parsedInput[0]
		switch strings.TrimSpace(command) {
		case PWD:
			p.execPwdCommand()
			continue
		case CD:
			if len(parsedInput) < 2 {
				continue
			}
			p.execCdCommand(parsedInput[1])
			continue
		case ECHO:
			if len(parsedInput) < 2 {
				continue
			}
			p.echo(parsedInput[1:])
			continue
		case CAT:
			if len(parsedInput) < 2 {
				continue
			}
			utils.ExecuteScript(parsedInput[0], parsedInput[1:]...)
			continue
		case TYPE:
			if len(parsedInput) < 2 {
				continue
			}
			p.execTypeCommand(parsedInput[1])
			continue
		}

		_, err := utils.FindFile(command)
		if err != nil {
			p.repl.Print(input, false)
			continue
		}

		utils.ExecuteScript(command, parsedInput[1:]...)
		continue
	}
}

func (p *shellProgram) echo(value []string) {
	output := strings.Join(value, " ")
	p.repl.Print(output, true)
}

func (p *shellProgram) execTypeCommand(arg string) {
	var output []string
	var outputString string

	exists := builtins.IsBuiltin(arg)
	if !exists {
		path, err := utils.FindFile(arg)
		if err != nil {
			p.repl.Print(arg, false)
			return
		}
		output = []string{arg, "is", path}
		outputString = utils.BuildStrings(output)
		p.repl.Print(outputString, true)
		return
	}

	output = []string{arg, "is a shell builtin"}
	outputString = utils.BuildStrings(output)
	p.repl.Print(outputString, true)
}

func (p *shellProgram) execPwdCommand() {
	output := utils.GetWorkingDirectory()
	p.repl.Print(output, true)
}

func (p *shellProgram) execCdCommand(path string) {
	err := utils.ChangeDirectory(path)
	if err != nil {
		p.repl.Print("cd: no such file or directory: "+path, true)
	}
}
