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

// Print to print the AST
func Print(e Expr) string {
	switch t := e.(type) {

	case Binary:
		exprs := []Expr{t.Left, t.Right}
		return parenthesize(t.Operator.Lexeme, exprs)

	case Grouping:
		exprs := []Expr{t.Expression}
		return parenthesize("group", exprs)

	case Literal:
		switch t.Value.(type) {
		case nil:
			return "nil"
		case string:
			s, _ := t.Value.(string)
			return s
		case float64, int, float32:
			f, _ := t.Value.(float64)
			return fmt.Sprintf("%f", f)
		default:
			return "error"
		}

	case Unary:
		exprs := []Expr{t.Right}
		return parenthesize(t.Operator.Lexeme, exprs)

	default:
		return "error"
	}
}

func parenthesize(name string, exprs []Expr) string {
	out := "(" + name

	for _, e := range exprs {
		out += " "
		out += Print(e)
	}

	out += ")"
	return out
}

func Test() {
	expression := Binary{
		Left: Unary{
			Operator: token.Token{
				Type:    token.MINUS,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right: Literal{Value: 123.0},
		},
		Operator: token.Token{
			Type:    token.STAR,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Right: Grouping{
			Expression: Literal{Value: 45.67},
		},
	}
	fmt.Println(Print(expression))
}
