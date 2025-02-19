package repl

type Repl interface {
	Read() (string, error)
	Print(string)
}

type Program interface {
	Run()
	exit()
	echo()
}
