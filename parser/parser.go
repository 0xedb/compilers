package parser

import (
	"github.com/thebashshell/compilers/ast"
	"github.com/thebashshell/compilers/lexer"
	"github.com/thebashshell/compilers/token"
)

type Parser struct {
	l         *lexer.Lexer
	cur, peek token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

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

	p.nextToken()

	if p.cur.Literal != token.IDENT {
		return nil
	}

	stmt.Variable = &ast.Identifier{Token: stmt.Token, Value: stmt.Token.Literal}

	// check if next is =
	if p.peek.Literal != token.ASSIGN {
		return nil
	}

	// consume expression until semi colon
	for p.cur.Literal != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}
