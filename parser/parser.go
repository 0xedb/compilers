package parser

import (
	"fmt"
	"strconv"

	"github.com/0xedb/compilers/ast"
	"github.com/0xedb/compilers/lexer"
	"github.com/0xedb/compilers/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peek.Type]; ok {
		return p
	}
	return LOWEST
}
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.cur.Type]; ok {
		return p
	}
	return LOWEST
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l         *lexer.Lexer
	cur, peek token.Token
	errors    []string

	infixes  map[token.TokenType]infixParseFn
	prefixes map[token.TokenType]prefixParseFn
}

func (p *Parser) nextToken() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l, errors: []string{},
		infixes:  map[token.TokenType]infixParseFn{},
		prefixes: map[token.TokenType]prefixParseFn{},
	}

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.cur,
		Operator: p.cur.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.cur,
		Operator: p.cur.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.cur}
	value, err := strconv.ParseInt(p.cur.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.cur.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.cur, Value: p.cur.Literal}
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixes[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixes[tokenType] = fn
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
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.cur}

	stmt.Expression = p.parseExpression(LOWEST)

	return nil
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixes[p.cur.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.cur.Type)
		return nil
	}
	left := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixes[p.peek.Type]
		if infix == nil {
			return left
		}
		p.nextToken()
		left = infix(left)
	}

	return left
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.cur}

	p.nextToken()

	// skip til semi
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
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
