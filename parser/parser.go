/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/

package parser

import (
	"GoLox/ast"
	"GoLox/lox_error"
	"GoLox/token"
	"errors"
)

// Parser is Lox parser
type Parser struct {
	tokens  []token.Token
	current int
}

// NewParser to create new parser
func NewParser(toks []token.Token) *Parser {
	return &Parser{
		tokens:  toks,
		current: 0,
	}
}

/* old Parser*/
// // Parse to Parse
// func (p *Parser) Parse() (ast.Expr, error) {
// 	exp, err := p.expression()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return exp, nil
// }

// program   → statement* EOF
// Parse to Parse
func (p *Parser) Parse() ([]ast.Stmt, error) {
	var statementsList []ast.Stmt

	for !p.isAtEnd() {
		statementsList = append(statementsList, p.statement()) // TODO: error
	}

	return statementsList, nil
}

// statement → exprStmt | printStmt
func (p *Parser) statement() (ast.Stmt, error) {
	if p.match(token.PRINT) {
		return p.printStatement()
	}

	return p.expressionStatement()
}

// printStmt → "print" expression ";"
func (p *Parser) printStatement() (ast.Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}

	conErr := p.consume(token.SEMICOLON, "Expect ';' after value.")
	if conErr != nil {
		return nil, conErr
	}

	return ast.PrintStatement{Value: value}, nil
}

//exprStmt  → expression ";"
func (p *Parser) expressionStatement() (ast.Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}

	conErr := p.consume(token.SEMICOLON, "Expect ';' after value.")
	if conErr != nil {
		return nil, conErr
	}

	return ast.ExpressionStatement{Expression: value}, nil
}

//   expression -> equality ;
func (p *Parser) expression() (ast.Expr, error) {
	return p.equality()
}

//   equality -> comparison ( ( "!=" | "==" ) comparison )*
func (p *Parser) equality() (ast.Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right, errR := p.comparison()
		if errR != nil {
			return nil, errR
		}
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

// comparison -> addition ( ( ">" | ">=" | "<" | "<=" ) addition )* ;
func (p *Parser) comparison() (ast.Expr, error) {
	expr, err := p.addition()
	if err != nil {
		return nil, err
	}

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right, errR := p.addition()
		if errR != nil {
			return nil, errR
		}
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

// addition -> multiplication ( ( "-" | "+" ) multiplication )* ;
func (p *Parser) addition() (ast.Expr, error) {
	expr, err := p.multiplication()
	if err != nil {
		return nil, err
	}

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right, errR := p.multiplication()
		if errR != nil {
			return nil, errR
		}
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

//multiplication -> unary ( ( "/" | "*" ) unary )* ;
func (p *Parser) multiplication() (ast.Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right, errR := p.unary()
		if errR != nil {
			return nil, errR
		}
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

/*
	unary -> ( "!" | "-" ) unary
            | primary ;
*/
func (p *Parser) unary() (ast.Expr, error) {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right, errR := p.unary()
		if errR != nil {
			return nil, errR
		}
		return ast.Unary{Operator: operator, Right: right}, nil
	}

	return p.primary()
}

// primary -> NUMBER | STRING | "false" | "true" | "nil" | "(" expression ")" ;
func (p *Parser) primary() (ast.Expr, error) {
	if p.match(token.FALSE) {
		return ast.Literal{Value: false}, nil
	}
	if p.match(token.TRUE) {
		return ast.Literal{Value: true}, nil
	}
	if p.match(token.NIL) {
		return ast.Literal{Value: nil}, nil
	}

	if p.match(token.NUMBER, token.STRING) {
		return ast.Literal{Value: p.previous().Literal}, nil
	}

	if p.match(token.LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		err = p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			return nil, err // TODO:
		}
		return ast.Grouping{Expression: expr}, nil
	}

	return nil, errors.New("Expect expression") // TODO:
}

// Helping functions
func (p *Parser) consume(tok token.Type, msg string) error {
	if p.check(tok) {
		p.advance()
		return nil
	}
	lox_error.AtToken(p.tokens[p.current], msg)
	return errors.New(msg)
}

func (p *Parser) match(toks ...token.Type) bool {
	for _, t := range toks {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

// TestParser is a func for testing
func TestParser() {
	// tok0 := token.Token{token.LEFT_PAREN, "(", nil, 1}
	// tok1 := token.Token{token.NUMBER, "4", 4.0, 1}
	// tok2 := token.Token{token.PLUS, "+", nil, 1}
	// tok3 := token.Token{token.STRING, "\"12\"", "\"12\"", 1}
	// tok4 := token.Token{token.RIGHT_PAREN, ")", nil, 1}
	// tok5 := token.Token{token.EOF, "", nil, 2}

	// tokens := []token.Token{tok0, tok1, tok2, tok3, tok4, tok5}

	// loxP := NewParser(tokens)

	// exp, _ := loxP.Parse()

	// fmt.Println(ast.Print(exp))
}
