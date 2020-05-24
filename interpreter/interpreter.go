/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/

package interpreter

import (
	"GoLox/ast"
	"GoLox/token"
	"errors"
	"fmt"
)

// Interpret is : The interpreter of lox
func Interpret(e ast.Expr) (interface{}, error) {
	value, err := evaluate(e)
	if err != nil {
		return nil, err
	}
	return value, nil
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
			return nil, err
		}

		switch t.Operator.Type {
		case token.MINUS:
			num, ok := right.(float64) // It's a number or string if string (ok = false)
			if !ok {
				// it won't hanppen
				msg := fmt.Sprint("The value (\"", right, "\") can not be cast to f64!!")
				return nil, errors.New(msg)
			}
			return -1 * num, nil

		case token.BANG:
			return !isTruthy(right), nil

		}
	}
	return nil, nil // will Never get to this
}

func evaluateBinary(b ast.Binary) (interface{}, error) {
	left, errL := evaluate(b.Left)
	if errL != nil {
		return nil, errL
	}

	right, errR := evaluate(b.Right)
	if errR != nil {
		return nil, errR
	}

	if b.Operator.Type == token.PLUS {

		switch L := left.(type) {
		case float64:
			R, ok := right.(float64)
			if !ok {
				msg := fmt.Sprint("right exp = (", right, ") is not numbers, the operation (", token.TypeNames[b.Operator.Type], ") can not be performed")
				return nil, errors.New(msg)
			}
			return L + R, nil

		case string:
			R, ok := right.(string)
			if !ok {
				msg := fmt.Sprint("right exp = (", right, ") is not string, the operation (", token.TypeNames[b.Operator.Type], ") can not be performed")
				return nil, errors.New(msg)
			}
			return L + R, nil

		}

	} else if b.Operator.Type == token.EQUAL_EQUAL {
		return left == right, nil

	} else if b.Operator.Type == token.BANG_EQUAL {
		return left != right, nil

	} else {
		numL, numR, err := checkNumberOperands(b.Operator.Type, left, right)
		if err != nil {
			return nil, err
		}
		switch b.Operator.Type {
		case token.MINUS:
			return numL - numR, nil

		case token.SLASH:
			return numL / numR, nil

		case token.STAR:
			return numL * numR, nil

		case token.GREATER:
			return numL > numR, nil

		case token.GREATER_EQUAL:
			return numL >= numR, nil

		case token.LESS:
			return numL < numR, nil

		case token.LESS_EQUAL:
			return numL <= numR, nil
		}
	}
	return nil, nil // will Never get to this
}

func checkNumberOperands(op token.Type, left, right interface{}) (float64, float64, error) {
	numL, okL := left.(float64)
	numR, okR := right.(float64)

	if !okL && !okR {
		msg := fmt.Sprint("Both (", left, ", ", right, ") are not numbers, the operation (", token.TypeNames[op], ") can not be performed")
		return 0.0, 0.0, errors.New(msg)
	} else if !okL {
		msg := fmt.Sprint("Left exp = (", left, ") is not numbers, the operation (", token.TypeNames[op], ") can not be performed")
		return 0.0, numR, errors.New(msg)
	} else if !okR {
		msg := fmt.Sprint("Right exp = (", right, ") is not numbers, the operation (", token.TypeNames[op], ")  can not be performed")
		return numL, 0.0, errors.New(msg)
	} else {
		return numL, numR, nil
	}
}

func isTruthy(obj interface{}) bool {
	switch t := obj.(type) {
	case nil:
		return false
	case bool:
		return t
	default:
		return true
	}
}

func Test() {
	tmp := ast.Binary{
		Left:     ast.Literal{Value: "float64(16)"},
		Operator: token.Token{Type: token.PLUS, Lexeme: "+", Literal: nil, Line: 1},
		Right:    ast.Literal{Value: "float64(24)"},
	}

	out, err := evaluateBinary(tmp)

	fmt.Println(out)
	fmt.Println(err)

}
