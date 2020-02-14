package token

import "fmt"

// TokenType represents a set of lexical tokens of the Lox programming language.
type TokenType int

// TokenType set elements, all the lexems for the language
const (
	EOF TokenType = iota

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

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal fmt.Stringer
	Line    int
}

func (tok *Token) toString() string {
	return fmt.Sprintf("Token [Type : %v, Lexeme : %s, Literal : %s, Line : %d]",
		tok.Type, tok.Lexeme, tok.Literal, tok.Line)
}
