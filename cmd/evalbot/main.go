package main

import (
	"fmt"

	"github.com/volokitty/evalbot/pkg/calculator"
)

func main() {
	result, calcErr := calculator.Calculate("3+4*2/(1-5)^2^3")

	if calcErr != nil {
		fmt.Println(calcErr)
	}

	fmt.Println(result)
}
