package interpreter

import (
	"GoLox/ast"
	"GoLox/token"
	"errors"
	"fmt"
)

func Interpret(e ast.Expr) {
	value, err := evaluate(e)
	if err != nil {
		// do something
	}
	fmt.Println(value)
}

func evaluate(e ast.Expr) (interface{}, error) {
	switch t := e.(type) {

	case ast.Binary:
		return evaluateBinary(t)

	case ast.Grouping:
		return evaluate(t.Expression)

	case ast.Literal:
		return t.Value, nil

	case ast.Unary:
		right, err := evaluate(t.Right)
		if err != nil {
			// do something
		}

		switch t.Operator.Type {
		case token.MINUS:
			num, ok := right.(float64)
			if !ok {
				return nil, errors.New("sksk")
			}
			return -1 * num, nil

		case token.BANG:
			return !isTruthy(right), nil

		}
		return nil, nil

	default:
		return "error", errors.New("sksk")
	}
}

func evaluateBinary(b ast.Binary) (interface{}, error) {
	left, errL := evaluate(b.Left)
	right, errR := evaluate(b.Right)

	if errL != nil || errR != nil {
		return nil, errors.New("sksk")
	}

	switch b.Operator.Type {
	case token.MINUS:
		numL, numR, err := checkNumberOperands(left, right)
		if err != nil {
			return nil, errors.New("TODO")
		}
		return numL - numR, nil

	case token.SLASH:
		numL, numR, err := checkNumberOperands(left, right)
		if err != nil {
			return nil, errors.New("TODO")
		}
		return numL / numR, nil

	case token.STAR:
		numL, numR, err := checkNumberOperands(left, right)
		if err != nil {
			return nil, errors.New("TODO")
		}
		return numL * numR, nil

	case token.PLUS:

	}
	return nil, nil
}

func checkNumberOperands(left, right interface{}) (float64, float64, error) {
	numL, okL := left.(float64)
	numR, okR := left.(float64)

	if okL && okR {
		return numL, numR, nil
	} else {
		return 0.0, 0.0, errors.New("TODO")
	}

}

func isTruthy(e ast.Expr) bool {
	if e == nil {
		return false
	}

	if t, ok := e.(bool); ok {
		return t
	}

	return true
}
