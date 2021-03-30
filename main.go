package main

import (
	"fmt"

	"github.com/thebashshell/compilers/ast"

	"github.com/thebashshell/compilers/lexer"
	"github.com/thebashshell/compilers/parser"
)

func main() {
	// repl.StartREPL(os.Stdin, os.Stdout)
	input := `let kofi = 39;`

	l := lexer.New(input)

	p := parser.New(l)
	fmt.Println(p.ParseProgram().Statements[0].(*ast.LetStatment).Variable)
}
