package repl

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
		p.repl.Print(input)
	}
}
