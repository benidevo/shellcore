package main

import "github.com/benidevo/shellcore/pkg/repl"

func main() {
	shellRepl := repl.NewShellRepl()
	program := repl.NewShellProgram(shellRepl)

	program.Run()
}
