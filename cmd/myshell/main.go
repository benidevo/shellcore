package main

import (
	"github.com/codecrafters-io/shell-starter-go/pkg/repl"
)

func runProgram() {
	repl := repl.NewShellRepl()
	for {
		input, _ := repl.Read()

		repl.Print(input)
	}
}

func main() {
	runProgram()
}
