package tools

import (
	"fmt"
)

// LoxGrammer is the grammer of lox
var LoxGrammer = map[string][]string{
	"Binary":   []string{"left", "operator", "right"},
	"Grouping": []string{"expression"},
	"Literal":  []string{"value"},
	"Unary":    []string{"operator", "right"},
}

// GenerateLoxAST to generate lox AST
func GenerateLoxAST(grammer map[string][]string, file string) {

}

// type Expr interface{}

// type Binary struct {
// 	left     Expr
// 	operator token.Token
// 	right    Expr
// }

// Test to test GenerateLoxAST
func Test() {
	fmt.Sprintf(`
	
	`)
}
