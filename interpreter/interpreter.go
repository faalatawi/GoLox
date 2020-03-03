package interpreter

import (
	"GoLox/ast"
	"fmt"
)

// LoxInterpreter is an empty struct
type LoxInterpreter struct{}

func Interpret(e ast.Expr) string {
	switch t := e.(type) {

	case ast.Binary:
		exprs := []ast.Expr{t.left, t.right}
		return parenthesize(t.operator.Lexeme, exprs)

	case ast.Grouping:
		exprs := []Expr{t.expression}
		return parenthesize("group", exprs)

	case ast.Literal:
		switch t.value.(type) {
		case nil:
			return "nil"
		case string:
			s, _ := t.value.(string)
			return s
		case float64, int, float32:
			f, _ := t.value.(float64)
			return fmt.Sprintf("%f", f)
		default:
			return "error"
		}

	case ast.Unary:
		exprs := []Expr{t.right}
		return parenthesize(t.operator.Lexeme, exprs)

	default:
		return "error"
	}
}
