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

// Token is a struct
type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

// ToString is a method
func (tok Token) ToString() string {
	return fmt.Sprintf("Token [Type : %v, Lexeme : %s, Literal : %v, Line : %d]",
		tok.Type, tok.Lexeme, tok.Literal, tok.Line)
}
