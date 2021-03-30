package repl

import (
	"bufio"
	"io"

	"github.com/thebashshell/compilers/token"

	"github.com/thebashshell/compilers/lexer"
)

func StartREPL(in io.Reader, out io.Writer) {
	const PROMPT = `>>> `
	const WELCOME = `Welcome to IntLang`
	const GREETING = `
	
	██▓ ███▄    █ ▄▄▄█████▓ ██▓    ▄▄▄       ███▄    █   ▄████ 
	▓██▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██▒   ▒████▄     ██ ▀█   █  ██▒ ▀█▒
	▒██▒▓██  ▀█ ██▒▒ ▓██░ ▒░▒██░   ▒██  ▀█▄  ▓██  ▀█ ██▒▒██░▄▄▄░
	░██░▓██▒  ▐▌██▒░ ▓██▓ ░ ▒██░   ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█  ██▓
	░██░▒██░   ▓██░  ▒██▒ ░ ░██████▒▓█   ▓██▒▒██░   ▓██░░▒▓███▀▒
	░▓  ░ ▒░   ▒ ▒   ▒ ░░   ░ ▒░▓  ░▒▒   ▓▒█░░ ▒░   ▒ ▒  ░▒   ▒ 
	 ▒ ░░ ░░   ░ ▒░    ░    ░ ░ ▒  ░ ▒   ▒▒ ░░ ░░   ░ ▒░  ░   ░ 
	 ▒ ░   ░   ░ ░   ░        ░ ░    ░   ▒      ░   ░ ░ ░ ░   ░ 
	 ░           ░              ░  ░     ░  ░         ░       ░ 
																

	`

	out.Write([]byte(GREETING))
	out.Write([]byte(WELCOME))
	out.Write([]byte("\n"))

	for {

		// print prompt
		out.Write([]byte(PROMPT))

		// get input
		scanner := bufio.NewScanner(in)

		if scanner.Scan() {
			l := lexer.New(scanner.Text())

			for output := l.NextToken(); output.Type != token.EOF; output = l.NextToken() {
				out.Write([]byte(output.String()))
			}
		}
	}

}
