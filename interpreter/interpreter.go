/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/

package interpreter

import (
	"GoLox/ast"
	"fmt"

	// loxErr "GoLox/error"
	"GoLox/token"
	"errors"
)

// Interpret is : The interpreter of lox
func Interpret(e ast.Expr) (interface{}, error) {
	value, err := evaluate(e)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// Check list:
// ast.Binary
// ast.Literal Done
// ast.Unary Done
// ast.Grouping Done

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

	if b.Operator.Type == token.PLUS { // TODO:
		numL, numR, err := checkNumberOperands(b.Operator.Type, left, right)
		if err != nil {
			return nil, err
		}
		return numL + numR, nil

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
		}
	}
	return nil, nil // will Never get to this
}

func checkNumberOperands(op token.Type, left, right interface{}) (float64, float64, error) {
	numL, okL := left.(float64)
	numR, okR := left.(float64)

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

}
