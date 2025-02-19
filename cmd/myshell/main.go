package main

import "github.com/codecrafters-io/shell-starter-go/pkg/repl"

func main() {
	shellRepl := repl.NewShellRepl()
	program := repl.NewShellProgram(shellRepl)

	program.Run()
}
