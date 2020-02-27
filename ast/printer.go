package ast

import (
	"GoLox/token"
	"fmt"
)

// Print to print the AST
func Print(e Expr) {
	fmt.Println(e.print()) // tree head
}

// ==============================================
// Implementing print method for all grammer

func (b Binary) print() string {
	exprs := []Expr{b.left, b.right}
	return parenthesize(b.operator.Lexeme, exprs)
}

func (g Grouping) print() string {
	exprs := []Expr{g.expression}
	return parenthesize("group", exprs)
}

func (l Literal) print() string {
	switch l.value.(type) {
	case nil:
		return "nil"
	case string:
		s, _ := l.value.(string)
		return s
	case float64, int, float32:
		f, _ := l.value.(float64)
		return fmt.Sprintf("%f", f)
	default:
		return "error"
	}
}

func (u Unary) print() string {
	exprs := []Expr{u.right}
	return parenthesize(u.operator.Lexeme, exprs)
}

func parenthesize(name string, exprs []Expr) string {
	out := "(" + name

	for _, e := range exprs {
		out += " "
		out += e.print()
	}

	out += ")"
	return out
}

func Test() {
	expression := Binary{
		left: Unary{
			operator: token.Token{
				Type:    token.MINUS,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			right: Literal{value: 123.0},
		},
		operator: token.Token{
			Type:    token.STAR,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		right: Grouping{
			expression: Literal{value: 45.67},
		},
	}

	Print(expression)
}
