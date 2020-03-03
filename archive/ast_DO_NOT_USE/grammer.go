package ast_DO_NOT_USE

import "GoLox/token"

type Expr interface {
	print() string // each visitor has a function in this interface
}

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
