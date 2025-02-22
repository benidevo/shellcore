package repl

type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Print(string, bool)
}

type Repl interface {
	Reader
	Writer
}

type Program interface {
	Run()
}
