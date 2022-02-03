package calculator

type associativity int

const (
	left associativity = iota
	right
)

type operator struct {
	precedence    int
	associativity associativity
}

var operators = map[string]operator{
	"^": {4, right},
	"*": {3, left},
	"/": {3, left},
	"+": {2, left},
	"-": {2, left},
}
