package parser

import "strings"

// ParseCommand splits the input string into a slice of arguments, handling quoted substrings.
//
// It iterates through each character of the input string, using a strings.Builder to accumulate characters
// for each argument. The function respects single quotes, treating the characters within quotes as a
// single argument, even if spaces are present. It removes surrounding quotes from arguments after parsing.
//
// Returns a slice of strings representing the parsed arguments.
func ParseCommand(input string) []string {
	var args []string
	var currentArg strings.Builder
	inQuotes := false

	for i := 0; i < len(input); i++ {
		char := input[i]
		switch char {
		case '\'':
			inQuotes = !inQuotes
		case ' ':
			if !inQuotes {
				if currentArg.Len() > 0 {
					args = append(args, currentArg.String())
					currentArg.Reset()
				}
			} else {
				currentArg.WriteByte(char)
			}
		case '"':
			if !inQuotes {
				if currentArg.Len() > 0 {
					args = append(args, currentArg.String())
					currentArg.Reset()
				}
			} else {
				currentArg.WriteByte(char)
			}
		default:
			currentArg.WriteByte(char)
		}
	}

	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	for i, arg := range args {
		if len(arg) > 0 && arg[0] == '\'' && arg[len(arg)-1] == '\'' {
			args[i] = arg[1 : len(arg)-1]
		}
	}
	for i, arg := range args {
		if len(arg) > 0 && arg[0] == '"' && arg[len(arg)-1] == '"' {
			args[i] = arg[1 : len(arg)-1]
		}
	}

	return args
}
