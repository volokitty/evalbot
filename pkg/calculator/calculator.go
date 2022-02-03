package calculator

func Calculate(infixExpr string) (string, error) {
	lexer := newLexer(infixExpr)
	tokens, lexErr := lexer.lex()

	if lexErr != nil {
		return "", lexErr
	}

	rpn, infixToRPNErr := infixToRPN(tokens)

	if infixToRPNErr != nil {
		return "", infixToRPNErr
	}

	result, evalErr := evalRPN(rpn)

	if evalErr != nil {
		return "", evalErr
	}

	return result, nil
}
