package ast2

import (
	"GoLox/token"
	"fmt"
)

// Print to print the AST
func Print(e Expr) string {
	switch e.(type) {
	case Binary:
		b := e.(Binary)
		exprs := []Expr{b.left, b.right}
		return parenthesize(b.operator.Lexeme, exprs)
	case Grouping:
		g := e.(Grouping)
		exprs := []Expr{g.expression}
		return parenthesize("group", exprs)
	case Literal:
		l := e.(Literal)
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
	case Unary:
		u := e.(Unary)
		exprs := []Expr{u.right}
		return parenthesize(u.operator.Lexeme, exprs)
	default:
		return "error"
	}
}

// ==============================================
// Implementing print method for all grammer

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

	fmt.Println(Print(expression))
}
