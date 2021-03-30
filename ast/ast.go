package ast

import (
	"github.com/thebashshell/compilers/token"
)

type Node interface {
	TokenLiteral() string
	// String() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	return "program"
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// func (i *Identifier) String() string       { return "identifier" }
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatment struct {
	Token    token.Token
	Variable *Identifier
	Value    Expression
}

func (l *LetStatment) statementNode() {}

// func (l *LetStatment) String() string       { return fmt.Sprintf("%s %s", l.Token, l.Variable) }
func (l *LetStatment) TokenLiteral() string { return l.Token.Literal }

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) statementNode()       {}
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }

type ExpressionStatement struct {
	Token token.Token
	Value Expression
}

func (r *ExpressionStatement) statementNode()       {}
func (r *ExpressionStatement) TokenLiteral() string { return r.Token.Literal }
