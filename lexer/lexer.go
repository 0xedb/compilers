package lexer

import (
	"github.com/thebashshell/compilers/token"
)

type Lexer struct {
	ch          byte
	pos, offset int
	input       string
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	return l
}

func (l *Lexer) readChar() {
	if l.offset >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.offset]
	}

	l.pos = l.offset
	l.offset++
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{}
	l.eatWhitespace()

	switch l.ch {
	case '=':
		// catches equal
		tok = token.MakeToken(token.EQ, l.ch)

	case '+':
		tok = token.MakeToken(token.PLUS, l.ch)

	case '-':
		tok = token.MakeToken(token.MINUS, l.ch)

	case '!':
		// catches not equl
		tok = token.MakeToken(token.BANG, l.ch)

	case '*':
		tok = token.MakeToken(token.ASTERISK, l.ch)

	case '/':
		tok = token.MakeToken(token.SLASH, l.ch)

	case '<':
		tok = token.MakeToken(token.LT, l.ch)

	case '>':
		tok = token.MakeToken(token.GT, l.ch)

	case ',':
		tok = token.MakeToken(token.COMMA, l.ch)

	case ';':
		tok = token.MakeToken(token.SEMICOLON, l.ch)

	case '(':
		tok = token.MakeToken(token.LPAREN, l.ch)

	case ')':
		tok = token.MakeToken(token.RPAREN, l.ch)

	case '{':
		tok = token.MakeToken(token.LBRACE, l.ch)

	case '}':
		tok = token.MakeToken(token.RBRACE, l.ch)

	case '[':
		tok = token.MakeToken(token.LBRAKET, l.ch)

	case ']':
		tok = token.MakeToken(token.RBRAKET, l.ch)

	case ':':
		tok = token.MakeToken(token.COLON, l.ch)

	case 0:
		return token.MakeToken(token.EOF, l.ch)

	default:
		// letters
		if token.IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if token.IsNumber(l.ch) {
			// numbers
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.MakeToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

// func (l *Lexer) nextChar() byte {

// }

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	pos := l.pos

	for token.IsNumber(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos

	for token.IsLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}
