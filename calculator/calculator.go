package calculator

import "fmt"

func Calculate(infixExpr string) (string, error) {
	lexer := newLexer(infixExpr)
	tokens, err := lexer.lex()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tokens)
	fmt.Println(infixToRPN(tokens))
	return "", nil
}
