package parser

import (
	"fmt"

	"github.com/thebashshell/compilers/ast"
	"github.com/thebashshell/compilers/lexer"
	"github.com/thebashshell/compilers/token"
)

type Parser struct {
	l         *lexer.Lexer
	cur, peek token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	program.Statements = []ast.Statement{}

	for p.cur.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.cur.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatment {
	stmt := &ast.LetStatment{Token: p.cur}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Variable = &ast.Identifier{Token: p.cur, Value: p.cur.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// consume expression until semi colon
	for p.cur.Literal != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.cur.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peek.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peek.Type)
	p.errors = append(p.errors, msg)
}
