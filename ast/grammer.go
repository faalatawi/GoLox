package ast

import "GoLox/token"

type Visitor interface {
	visitBinary(*Binary) interface{}
	visitGrouping(*Grouping) interface{}
	visitLiteral(*Literal) interface{}
	visitUnary(*Unary) interface{}
}

// Expr accept a visitor
type Expr interface { // accept a visitor
	Accept(Visitor) interface{}
}

type Binary struct {
	left     Expr
	operator token.Token
	right    Expr
}

func (b *Binary) Accept(v Visitor) interface{} {
	return v.visitBinary(b)
}

type Grouping struct {
	expression Expr
}

func (g *Grouping) Accept(v Visitor) interface{} {
	return v.visitGrouping(g)
}

type Literal struct {
	value interface{}
}

func (l *Literal) Accept(v Visitor) interface{} {
	return v.visitLiteral(l)
}

type Unary struct {
	operator token.Token
	right    Expr
}

func (u *Unary) Accept(v Visitor) interface{} {
	return v.visitUnary(u)
}
