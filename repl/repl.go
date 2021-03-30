package repl

import "io"

func StartREPL(in io.Reader, out io.Writer) {
	const PROMPT = `>>>`
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
}