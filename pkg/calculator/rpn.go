package calculator

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func infixToRPN(tokens []string) ([]string, error) {
	var rpnTokens []string
	operatorStack := stack.New()

	for _, token := range tokens {
		_, operator := operators[token]
		switch {
		case operator:
			for o1, o2 := token, operatorStack.Peek(); (o2 != nil && o2 != "(") && ((operators[o2.(string)].precedence > operators[o1].precedence) || (operators[o2.(string)].precedence == operators[o1].precedence && operators[o1].associativity == left)); o2 = operatorStack.Peek() {
				rpnTokens = append(rpnTokens, operatorStack.Pop().(string))
			}
			operatorStack.Push(token)
		case token == "(":
			operatorStack.Push(token)
		case token == ")":
			for operatorStack.Peek().(string) != "(" {
				rpnTokens = append(rpnTokens, operatorStack.Pop().(string))

				if operatorStack.Len() == 0 {
					return nil, fmt.Errorf("no matched parenthesis")
				}
			}
			operatorStack.Pop()
		default:
			rpnTokens = append(rpnTokens, token)
		}
	}

	for operatorStack.Len() != 0 {
		if operatorStack.Peek() == "(" {
			return nil, fmt.Errorf("no matched parenthesis")
		}

		rpnTokens = append(rpnTokens, operatorStack.Pop().(string))
	}

	return rpnTokens, nil
}
