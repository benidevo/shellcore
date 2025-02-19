package utils

import (
	"os/exec"
	"strings"
)

// FindFile finds the path of a given executable in the system's PATH.
//
// It uses the exec.LookPath function to search for the given filename in the
// system's PATH. If the file is found, it returns the path of the executable.
// If the file is not found, it returns an empty string and the error
// encountered.
func FindFile(filename string) (string, error) {
	path, err := exec.LookPath(filename)
	if err != nil {
		return "", err
	}
	return path, nil
}

// BuildStrings takes a slice of strings and builds a single string out of them by
// concatenating the strings with a space in between. It uses a strings.Builder to
// avoid extra allocations.
func BuildStrings(args []string) string {
	var response strings.Builder
	for _, arg := range args {
		response.WriteString(arg)
		response.WriteString(" ")
	}
	return response.String()
}
