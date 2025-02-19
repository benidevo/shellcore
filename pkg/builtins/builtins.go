package builtins

var builtinsMap = map[string]struct{}{
	"echo": {},
}

// IsBuiltin checks if the given command is a shell builtin.
// It returns true if the command exists in the builtinsMap, otherwise false.
func IsBuiltin(command string) bool {
	_, exists := builtinsMap[command]
	return exists
}
