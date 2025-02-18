package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v", err)
	}
	fmt.Printf("invalid command: %s", input)
}
