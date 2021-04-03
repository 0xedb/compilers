package main

import (
	"os"

	"github.com/0xedb/compilers/repl"
)

func main() {
	repl.StartREPL(os.Stdin, os.Stdout)
}
