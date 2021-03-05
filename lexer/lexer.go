package lexer

import "github.com/thebashshell/compilers/lexer/token"

// Lexer represents a lexer
type Lexer struct {
	Input        string
	Position     int
	ReadPosition int
	Ch           byte
}

// New returns a new lexer
func New(input string) *Lexer {
	l := &Lexer{Input: input}
	l.ReadChar()

	return l
}

// ReadChar reads and advances position
func (l *Lexer) ReadChar() {
	if l.ReadPosition >= len(l.Input) {
		l.Ch = 0
		return
	}

	l.Ch = l.Input[l.ReadPosition]
	l.Position = l.ReadPosition

	l.ReadPosition++
}

// NextToken returns the next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.Ch {
	case '=':
		tok = NewToken(token.ASSIGN, l.Ch)
	case ';':
		tok = NewToken(token.SEMICOLON, l.Ch)
	case '(':
		tok = NewToken(token.LPAREN, l.Ch)
	case ')':
		tok = NewToken(token.RPAREN, l.Ch)
	case ',':
		tok = NewToken(token.COMMA, l.Ch)
	case '+':
		tok = NewToken(token.PLUS, l.Ch)
	case '{':
		tok = NewToken(token.LBRACE, l.Ch)
	case '}':
		tok = NewToken(token.RBRACE, l.Ch)
	case 0:
		tok.Literal = ""
		tok.Kind = token.EOF
	default:
		if IsLetter(l.Ch) {
			tok.Literal = l.ReadIdentifier()
			return tok
		} else {
			return NewToken(token.ILLEGAL, l.Ch)
		}

	}

	l.ReadChar()
	return tok
}

// NewToken generates a new token
func NewToken(tk token.Type, ch byte) token.Token {
	return token.Token{Kind: tk, Literal: string(ch)}
}

// ReadIdentifier reads an identifier
func (l *Lexer) ReadIdentifier() string {
	position := l.Position

	for IsLetter(l.Ch) {
		l.ReadChar()
	}

	return l.Input[position:l.Position]
}

// IsLetter tells if ch is a letter
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
 