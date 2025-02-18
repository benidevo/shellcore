package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v", err)
	}
	input = strings.TrimSuffix(input, "\n")
	fmt.Printf("%s: invalid command", input)
}
