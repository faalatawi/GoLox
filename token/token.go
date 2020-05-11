/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/
package token

import "fmt"

// Type represents a set of lexical tokens of the Lox programming language.
type Type int

// TokenType set elements, all the lexems for the language
const (
	EOF Type = iota

	// single characters
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	SEMICOLON

	// single character operators
	BANG
	EQUAL
	GREATER
	LESS
	MINUS
	PLUS
	SLASH
	STAR

	// double character operators
	BANG_EQUAL
	EQUAL_EQUAL
	GREATER_EQUAL
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
)

var TypeNames = []string{
	"EOF",
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"COMMA",
	"DOT",
	"SEMICOLON",

	"BANG",
	"EQUAL",
	"GREATER",
	"LESS",
	"MINUS",
	"PLUS",
	"SLASH",
	"STAR",

	"BANG_EQUAL",
	"EQUAL_EQUAL",
	"GREATER_EQUAL",
	"LESS_EQUAL",

	"IDENTIFIER",
	"STRING",
	"NUMBER",

	"AND",
	"CLASS",
	"ELSE",
	"FALSE",
	"FUN",
	"FOR",
	"IF",
	"NIL",
	"OR",
	"PRINT",
	"RETURN",
	"SUPER",
	"THIS",
	"TRUE",
	"VAR",
	"WHILE",
}

// Token is a struct
type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

// ToString is a method
func (tok Token) ToString() string {
	return fmt.Sprintf("Token [Type : %s, Lexeme : %s, Literal : %v, Line : %d]",
		TypeNames[tok.Type], tok.Lexeme, tok.Literal, tok.Line)
}
