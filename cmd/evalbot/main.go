package main

import "github.com/volokitty/evalbot/calculator"

func main() {
	calculator.Calculate("3+4*2/(1-5)^2^3")
	calculator.Calculate("3+5*2")
	calculator.Calculate("4 + 4 * 2 / ( 1 - 5 )")
}
