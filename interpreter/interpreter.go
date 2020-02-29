package interpreter

import (
	"GoLox/ast"
	"fmt"
)

// LoxInterpreter is an empty struct
type LoxInterpreter struct{}

func (inter LoxInterpreter) Interpret(e ast.Expr) {
	value := e.Accept(inter)
	fmt.Println(value)
}

func (inter LoxInterpreter) visitBinary(b ast.Binary) interface{} {
	return nil
}

func (inter LoxInterpreter) visitGrouping(g ast.Grouping) interface{} {
	return nil
}

func (inter LoxInterpreter) visitLiteral(l ast.Literal) interface{} {
	return nil
}

func (inter LoxInterpreter) visitUnary(u ast.Unary) interface{} {
	return nil
}
