package scanner

import (
	"GoLox/token"
	"fmt"
)

// Scanner is an struct
type Scanner struct {
	source  string        // source code
	Tokens  []token.Token // tokens TODO: add * if needed
	current int
	start   int
	line    int
}

// Keywords of lox
var Keywords = map[string]token.Type{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

// NewScanner create new scanner
func NewScanner(source string) *Scanner {
	loxScan := Scanner{
		source:  source,
		Tokens:  make([]token.Token, 0),
		current: 0,
		start:   0,
		line:    1,
	}
	return &loxScan
}

// func (scan *Scanner) ScanTokens() []token.Token {
// 	for !scan.isAtEnd(){
// 		scan.start = scan.current
// 		scan.scanToken()
// 	}

// }

func (scan *Scanner) advance() string {
	c := scan.source[scan.current]
	scan.current++
	return string(c)
}

// Test for testing
func Test() {
	scan := NewScanner("abc")

	fmt.Println(scan)

	c := scan.advance()

	fmt.Println("c:", c)
	fmt.Println(scan)

}
