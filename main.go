package main

import (
	"os"

	"github.com/thebashshell/compilers/repl"
)

func main() {
	repl.StartREPL(os.Stdin, os.Stdout)
}
