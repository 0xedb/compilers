package token

// Type is a token type
type Type string 


// Token represents a token
type Token struct {
	Kind Token
	Literal string	
}