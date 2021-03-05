package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	
	// Identifiers + literals
	
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	
	// Operators
	
	ASSIGN = "="
	PLUS   = "+"
	
	// Delimiters

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	
	
	// Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Type is a token type
type Type string

// Token represents a token
type Token struct {
	Kind    Type
	Literal string
}
