package ast2

import "GoLox/token"

type Expr interface{}

type Binary struct {
	left     Expr
	operator token.Token
	right    Expr
}

type Grouping struct {
	expression Expr
}

type Literal struct {
	value interface{}
}

type Unary struct {
	operator token.Token
	right    Expr
}
