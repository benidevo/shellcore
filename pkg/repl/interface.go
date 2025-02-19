package repl

type Repl interface {
	Read() (string, error)
	Print(string, bool)
}

type Program interface {
	Run()
	exit()
	echo()
}
