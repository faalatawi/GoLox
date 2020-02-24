package main

import (
	"GoLox/scanner"
	"fmt"
	"go/token"
)

func main() {
	// myTok := token.Token{
	// 	Type:    token.CLASS,
	// 	Lexeme:  "class",
	// 	Literal: "fafa",
	// 	Line:    0,
	// }
	// fmt.Println(myTok.ToString())
	scanner.Test()

	fmt.Println(token.EOF)
}
