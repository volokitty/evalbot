package calculator

import (
	"fmt"
	"strconv"

	"github.com/golang-collections/collections/stack"
	"github.com/volokitty/evalbot/pkg/utils"
)

func evalRPN(tokens []string) (string, error) {
	evalStack := stack.New()
	for _, token := range tokens {
		evalStack.Push(token)
		if _, isOperator := operators[token]; isOperator {
			var result int

			operator := evalStack.Pop().(string)

			b, _ := strconv.Atoi(evalStack.Pop().(string))

			if evalStack.Peek() == nil {
				return "", fmt.Errorf("extra arithmetic operations")
			}

			a, _ := strconv.Atoi(evalStack.Pop().(string))

			switch operator {
			case "^":
				result = utils.PowInt(a, b)
			case "*":
				result = a * b
			case "/":
				result = a / b
			case "+":
				result = a + b
			case "-":
				result = a - b
			}

			evalStack.Push(strconv.Itoa(result))
		}
	}

	return evalStack.Pop().(string), nil
}
