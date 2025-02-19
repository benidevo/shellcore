package repl

type Repl interface {
	Read() (string, error)
	Evaluate(string) (string, error)
	Print(string)
}
