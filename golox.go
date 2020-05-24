/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/

package main

import (
	"GoLox/ast"
	"GoLox/interpreter"
	"GoLox/parser"
	"GoLox/scanner"
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println("Usage: golox [script]")
	} else if len(args) == 2 {
		// file_name := args[1]
		// runFile(file_name)
		panic("Not implemented yet")
	} else {
		runPrompt()
	}
}

func runPrompt() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		input.Scan()
		line := input.Text()

		if line == "exit" {
			return
		}

		lox_scanner := scanner.NewScanner(line)
		toks := lox_scanner.ScanTokens()

		scanner.PrintTokenList(toks)

		lox_parser := parser.New(toks)

		exp, err := lox_parser.Parse()
		if err != nil {
			fmt.Println(err)
			panic("error")
		}

		fmt.Println(ast.Print(exp))

		value, errInter := interpreter.Interpret(exp)
		if errInter != nil {
			fmt.Println(errInter)
			panic("error")
		}

		fmt.Println("==> value: ", value, "\n\n")

	}
}
