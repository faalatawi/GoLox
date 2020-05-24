/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/

package scanner

import (
	loxErr "GoLox/error"
	"GoLox/token"
	"fmt"
	"strconv"
)

// Scanner is an struct
type Scanner struct {
	source  string        // source code
	Tokens  []token.Token // tokens
	current int           // current index
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

// PunctuationMap is a map of lox Punctuation
var PunctuationMap = map[string]token.Type{
	"{":  token.LEFT_BRACE,
	"}":  token.RIGHT_BRACE,
	"(":  token.LEFT_PAREN,
	")":  token.RIGHT_PAREN,
	",":  token.COMMA,
	".":  token.DOT,
	"-":  token.MINUS,
	"+":  token.PLUS,
	";":  token.SEMICOLON,
	"*":  token.STAR,
	"!":  token.BANG,
	"!=": token.BANG_EQUAL,
	"=":  token.EQUAL,
	"==": token.EQUAL_EQUAL,
	"<":  token.LESS,
	"<=": token.LESS_EQUAL,
	">":  token.GREATER,
	">=": token.GREATER_EQUAL,
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

// ScanTokens to scan tokens and return an array of tokens
func (scan *Scanner) ScanTokens() []token.Token {
	for !scan.isAtEnd() {
		scan.start = scan.current
		scan.scanToken()
	}

	t := token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: scan.line}

	scan.Tokens = append(scan.Tokens, t)

	return scan.Tokens
}

func (scan *Scanner) scanToken() {
	c := scan.advance()

	switch c {
	case "{", "}", "(", ")", ";", ".", ",", "+", "-", "*":
		scan.addToken(PunctuationMap[c], nil)

	case "!", "=", "<", ">":
		if scan.match("=") {
			c += "="
		}
		scan.addToken(PunctuationMap[c], nil)

	case "/":
		if scan.match("/") {
			for scan.peek() != "\n" && !scan.isAtEnd() {
				scan.advance()
			}
		} else {
			scan.addToken(token.SLASH, nil)
		}

	case " ", "\t", "\r":
		break

	case "\n":
		scan.line++

	case "\"":
		scan.string()

	default:
		if isDigit(c) {
			scan.number()
		} else if isAlpha(c) {
			scan.identifier()
		} else {
			message := fmt.Sprintf("Unexpected character. :  %s", c)
			loxErr.AtLine(scan.line, message)
		}
	}
}

func (scan *Scanner) identifier() {
	for isAlphaNumeric(scan.peek()) {
		scan.advance()
	}

	start := scan.start
	end := scan.current
	key := scan.source[start:end]

	if val, ok := Keywords[key]; ok {
		scan.addToken(val, key)
	} else {
		scan.addToken(token.IDENTIFIER, key)
	}
}

func (scan *Scanner) number() {
	for isDigit(scan.peek()) {
		scan.advance()
	}

	if scan.peek() == "." && isDigit(scan.peekNext()) {
		// Consume the "."
		scan.advance()
	}

	for isDigit(scan.peek()) {
		scan.advance()
	}

	start := scan.start
	end := scan.current
	StrVal := scan.source[start:end]
	FloVal, _ := strconv.ParseFloat(StrVal, 64)
	scan.addToken(token.NUMBER, FloVal)
}

func (scan *Scanner) peekNext() string {
	next := scan.current + 1
	if next >= len(scan.source) {
		return "\000"
	}
	return string(scan.source[next])
}

func isAlphaNumeric(c string) bool {
	return isDigit(c) || isAlpha(c)
}

func isDigit(ch string) bool {
	tmp := []rune(ch)
	c := tmp[0]
	return '0' <= c && c <= '9'
}

func isAlpha(ch string) bool {
	tmp := []rune(ch)
	c := tmp[0]
	az := 'a' <= c && c <= 'z'
	AZ := 'A' <= c && c <= 'Z'
	return az || AZ || (c == '_')
}

func (scan *Scanner) string() {
	for scan.peek() != "\"" && !scan.isAtEnd() {
		if scan.peek() == "\n" {
			scan.line++
		}
		scan.advance()
	}

	//Unterminated string.
	if scan.isAtEnd() {
		loxErr.AtLine(scan.line, "Unterminated string.") // TODO ERROR
		return
	}

	// The closing ".
	scan.advance()

	start := scan.start + 1
	end := scan.current - 1
	value := scan.source[start:end]

	scan.addToken(token.STRING, value)
}

func (scan Scanner) peek() string {
	if scan.isAtEnd() {
		return "\000"
	}
	return string(scan.source[scan.current])
}

func (scan *Scanner) match(expected string) bool {
	if scan.isAtEnd() {
		return false
	} else if string(scan.source[scan.current]) != expected {
		return false
	} else {
		scan.current++
		return true
	}
}

func (scan *Scanner) addToken(tokType token.Type, literal interface{}) {
	lexeme := scan.source[scan.start:scan.current]
	t := token.Token{Type: tokType, Lexeme: lexeme, Literal: literal, Line: scan.line}
	scan.Tokens = append(scan.Tokens, t)
}

func (scan Scanner) isAtEnd() bool {
	return scan.current >= len(scan.source)
}

func (scan *Scanner) advance() string {
	c := scan.source[scan.current]
	scan.current++
	return string(c)
}

// PrintTokenList is a func to print the list
func PrintTokenList(toks []token.Token) {
	for _, t := range toks {
		fmt.Println(t.ToString())
	}
}

// Test for testing the Scanner
func Test() {
	source := `
	    var x = 12.1
	    if else
	    for
	    // kdjkdkkd
	    /
	    {}
	    ()
	    print
	    "fias'' // "

	    class
	    student
	    
	    // this is a comment
	(( )){} // grouping stuff
	!*+-/=<> <= == // operators
		`
	scan := NewScanner(source)

	tList := scan.ScanTokens()
	for _, t := range tList {
		fmt.Println(t.ToString())
	}
}
