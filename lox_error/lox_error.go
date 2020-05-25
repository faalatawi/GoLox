/*
 Copyright (c) 2020 Faisal Alatawi. All rights reserved
 Using this source code is governed by an MIT license
 you can find it in the LICENSE file.
*/

package lox_error

import (
	"GoLox/token"
	"fmt"
)

func report(line int, where, message string) {
	text := fmt.Sprintf("[line %d ] Error %s :  %s ", line, where, message)
	fmt.Println(text)
}

// AtLine to report an error in a line
func AtLine(line int, message string) {
	report(line, "", message)
}

// AtToken to report an error in a token
func AtToken(tok token.Token, message string) {
	if tok.Type == token.EOF {
		report(tok.Line, "  at end", message)

	} else {
		where := fmt.Sprintf("at '%v' ", tok.Lexeme)
		report(tok.Line, where, message)
	}
}
