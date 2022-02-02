package calculator

import (
	"fmt"
	"unicode"
)

type state int

const (
	start state = iota
	whitespace
	digit
	number
	lpar
	rpar
	op
	unknown
)

type lexer struct {
	expr      string
	state     state
	nextState state
}

func newLexer(expr string) *lexer {
	return &lexer{expr, start, -1}
}

func (l *lexer) getNextState(r rune) state {
	switch l.state {
	case start:
		_, operator := operators[string(r)]
		switch {
		case unicode.IsSpace(r):
			return whitespace
		case unicode.IsDigit(r):
			return digit
		case operator:
			return op
		case r == '(':
			return lpar
		case r == ')':
			return rpar
		default:
			return unknown
		}
	case whitespace:
		return start
	case op:
		return start
	case digit:
		if unicode.IsDigit(r) {
			return digit
		}
		return number
	case lpar:
		return start
	case rpar:
		return start
	case number:
		return start
	default:
		return unknown
	}
}

func (l *lexer) lex() ([]string, error) {
	var tokens []string
	var token string
	pos := 0

	for pos < len(l.expr) || l.nextState == unknown {
		l.nextState = l.getNextState(rune(l.expr[pos]))

		if l.nextState == start {
			if l.state != whitespace {
				tokens = append(tokens, token)
			}

			token = ""
		}

		if l.state == start || l.state == l.nextState {
			token += string(l.expr[pos])

			if pos == len(l.expr)-1 {
				tokens = append(tokens, token)
			}

			pos++
		}

		l.state = l.nextState
	}

	if l.nextState == unknown {
		pos--
		return tokens, fmt.Errorf("unknown character: %c", rune(l.expr[pos]))
	}

	return tokens, nil
}
