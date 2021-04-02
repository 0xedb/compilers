package repl

import (
	"bufio"
	"io"

	"github.com/0xedb/compilers/parser"

	"github.com/0xedb/compilers/lexer"
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
			p := parser.New(l)

			pg := p.ParseProgram()

			if len(p.Errors()) != 0 {
				printParserErrors(out, p.Errors())
				continue
			}

			// for output := l.NextToken(); output.Type != token.EOF; output = l.NextToken() {
			// 	out.Write([]byte(output.String()))
			// }

			io.WriteString(out, pg.String())
			io.WriteString(out, "\n")
		}
	}

}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
