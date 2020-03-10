/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/

package ast

import "GoLox/token"

type Expr interface{}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

type Grouping struct {
	Expression Expr
}

type Literal struct {
	Value interface{}
}

type Unary struct {
	Operator token.Token
	Right    Expr
}
