package calculator

import "fmt"

func Calculate(infixExpr string) (string, error) {
	lexer := newLexer(infixExpr)
	tokens, _ := lexer.lex()
	fmt.Println(tokens)
	return "", nil
}
