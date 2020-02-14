package main

import (
	"GoLox/token"
	"fmt"
)

type Number float64

func (n Number) String() string {
 	return fmt.Sprint(float64(n))
}

func main() {
	fmt.Println("Faisal")

	myTok := &token.Token{
		Type:    token.CLASS,
		Lexeme:  "class",
		Literal: Number(12),
		Line:    0,
	}

	fmt.Println(myTok.toString())

}
