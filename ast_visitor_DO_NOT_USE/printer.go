package ast_visitor_DO_NOT_USE

import (
	"GoLox/token"
	"fmt"
)

//ASTPrinter is  empty visitor
type ASTPrinter struct{}

func (v ASTPrinter) Print(e Expr) string {
	s, _ := e.Accept(v).(string)
	return s
}

// ==============================================
// Implementing print method for all grammer

func (v ASTPrinter) visitBinary(b *Binary) interface{} {
	exprs := []Expr{b.left, b.right}
	return v.parenthesize(b.operator.Lexeme, exprs)
}

func (v ASTPrinter) visitGrouping(g *Grouping) interface{} {
	exprs := []Expr{g.expression}
	return v.parenthesize("group", exprs)
}

func (v ASTPrinter) visitLiteral(l *Literal) interface{} {
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

func (v ASTPrinter) visitUnary(u *Unary) interface{} {
	exprs := []Expr{u.right}
	return v.parenthesize(u.operator.Lexeme, exprs)
}

func (v ASTPrinter) parenthesize(name string, exprs []Expr) string {
	out := "(" + name

	for _, e := range exprs {
		out += " "
		out += v.Print(e)
	}

	out += ")"
	return out
}

func Test() {
	expression := &Binary{
		left: &Unary{
			operator: token.Token{
				Type:    token.MINUS,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			right: &Literal{value: 123.0},
		},
		operator: token.Token{
			Type:    token.STAR,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		right: &Grouping{
			expression: &Literal{value: 45.67},
		},
	}

	printer := ASTPrinter{}

	fmt.Println(printer.Print(expression))

}
