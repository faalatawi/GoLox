package parser

import (
	"GoLox/ast"
	"GoLox/token"
	"errors"
)

type Parser struct {
	tokens  []token.Token
	current int
}

func New(toks []token.Token) Parser {
	return Parser{
		tokens:  toks,
		current: 0,
	}
}

func (p Parser) Parse() ast.Expr {
	exp, err := p.expression()
	if err != nil {
		// TODO:
		return nil
	}
	return exp
}

// 1)  expression -> equality ;
func (p Parser) expression() (ast.Expr, error) {
	return p.equality()
}

// 2)  equality -> comparison ( ( "!=" | "==" ) comparison )*
func (p Parser) equality() (ast.Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, errors.New(" ") // TODO:
	}

	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right, errR := p.comparison()
		if errR != nil {
			return nil, errors.New(" ") // TODO:
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
func (p Parser) comparison() (ast.Expr, error) {
	expr, err := p.addition()
	if err != nil {
		return nil, errors.New(" ") // TODO:
	}

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right, errR := p.addition()
		if errR != nil {
			return nil, errors.New(" ") // TODO:
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
func (p Parser) addition() (ast.Expr, error) {
	expr, err := p.multiplication()
	if err != nil {
		return nil, errors.New(" ") // TODO:
	}

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right, errR := p.multiplication()
		if errR != nil {
			return nil, errors.New(" ") // TODO:
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
func (p Parser) multiplication() (ast.Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, errors.New(" ") // TODO:
	}

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right, errR := p.unary()
		if errR != nil {
			return nil, errors.New(" ") // TODO:
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
func (p Parser) unary() (ast.Expr, error) {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right, errR := p.unary()
		if errR != nil {
			return nil, errors.New(" ") // TODO:
		}
		return ast.Unary{
			Operator: operator,
			Right:    right,
		}, nil
	}

	return p.primary()
}

// primary -> NUMBER | STRING | "false" | "true" | "nil" | "(" expression ")" ;
func (p Parser) primary() (ast.Expr, error) {
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
			return nil, errors.New(" ") // TODO:
		}
		err = p.consume(token.RIGHT_PAREN)
		if err != nil {
			return nil, errors.New("Expect ')' after expression.") // TODO:
		}
		return ast.Grouping{Expression: expr}, nil
	}

	return nil, errors.New("expect expression")
}

// Helping functions
func (p Parser) match(toks ...token.Type) bool {
	for _, t := range toks {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p Parser) check(t token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p Parser) previous() token.Token {
	return p.tokens[p.current-1]
}
