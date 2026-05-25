package parser

import (
	"fmt"
	"tree-walk-interpreter/lox"
	"tree-walk-interpreter/parser/grammar"
	"tree-walk-interpreter/parser/statement"
	"tree-walk-interpreter/token"
)

type Parser struct {
	Tokens  []token.Token
	Current int
}

type ParseError struct {
	Message string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse error: %s", e.Message)
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{Tokens: tokens, Current: 0}
}

func (p *Parser) Parse() ([]statement.Stmt, error) {
	statements := []statement.Stmt{}
	for !p.isAtEnd() {
		stmt, err := p.statement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}
	return statements, nil
}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, typ := range types {
		if p.check(typ) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.Current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() token.Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) check(typ token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == typ
}

func (p *Parser) consume(typ token.TokenType, message string) (token.Token, error) {
	if p.check(typ) {
		return p.advance(), nil
	}

	return p.peek(), p.error(p.peek(), message)
}

func (p *Parser) error(token token.Token, message string) ParseError {
	lox.Error(token.Line, message)
	return ParseError{Message: message}
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}
		switch p.peek().Type {
		case token.CLASS, token.FUN, token.VAR, token.FOR, token.IF, token.WHILE, token.PRINT, token.RETURN:
			return
		}
		p.advance()
	}
}

func (p *Parser) statement() (statement.Stmt, error) {
	if p.match(token.PRINT) {
		return p.printStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() (statement.Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(token.SEMICOLON, "Expect ';' after value.")
	if err != nil {
		return nil, err
	}
	return statement.NewPrintStmt(value), nil
}

func (p *Parser) expressionStatement() (statement.Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(token.SEMICOLON, "Expect ';' after expression.")
	if err != nil {
		return nil, err
	}
	return statement.NewExpressionStmt(expr), nil
}

func (p *Parser) expression() (grammar.Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) equality() (grammar.Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = grammar.NewBinary(expr, operator, right)
	}
	return expr, nil
}

func (p *Parser) comparison() (grammar.Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = grammar.NewBinary(expr, operator, right)
	}
	return expr, nil
}

func (p *Parser) term() (grammar.Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = grammar.NewBinary(expr, operator, right)
	}
	return expr, nil
}

func (p *Parser) factor() (grammar.Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = grammar.NewBinary(expr, operator, right)
	}
	return expr, nil
}

func (p *Parser) unary() (grammar.Expr, error) {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return grammar.NewUnary(operator, right), nil
	}

	return p.primary()
}

func (p *Parser) primary() (grammar.Expr, error) {
	if p.match(token.FALSE) {
		return grammar.NewLiteral(false), nil
	}
	if p.match(token.TRUE) {
		return grammar.NewLiteral(true), nil
	}
	if p.match(token.NIL) {
		return grammar.NewLiteral(nil), nil
	}

	if p.match(token.NUMBER, token.STRING) {
		return grammar.NewLiteral(p.previous().Literal), nil
	}

	if p.match(token.LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return grammar.NewGrouping(expr), nil
	}

	return nil, p.error(p.peek(), "Expect expression.")
}
