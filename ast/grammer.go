/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/

package ast

import (
	"GoLox/token"
	"fmt"
)

type Expr interface {
	isExpr() bool
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (b Binary) isExpr() bool { return true }

type Grouping struct {
	Expression Expr
}

func (g Grouping) isExpr() bool { return true }

type Literal struct {
	Value interface{}
}

func (l Literal) isExpr() bool { return true }

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (u Unary) isExpr() bool { return true }

// ==================================================

type Stmt interface {
	isStmt() bool
}

type PrintStatement struct {
	Value Expr
}

func (p PrintStatement) isStmt() bool { return true }

type ExpressionStatement struct {
	Expression Expr
}

func (e ExpressionStatement) isStmt() bool { return true }

func TestGrammer() {
	var lis []Stmt

	lis = append(lis, PrintStatement{Literal{123}})

	fmt.Println(lis)
}
