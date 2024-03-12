package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

var operators = map[string]int{
	"^": 1,
	"*": 2,
	"/": 2,
	"+": 3,
	"-": 3,
}

func PostfixToInfix(input string) (string, error) {
	tokens := strings.Split(input, " ")
	if len(tokens) == 0 {
		return "", nil
	}

	var stack []string
	prevPrecedence := 0
	for _, token := range tokens {
		if precedence, isOperator := operators[token]; isOperator {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid postfix expression")
			}

			operator := token
			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			if isNegative := strings.HasPrefix(operand2, "-"); isNegative {
				if (operator) == "+" {
					operator = "-"
					operand2 = operand2[1:]
				} else if (operator) == "-" {
					operator = "+"
					operand2 = operand2[1:]
				}
			}

			if prevPrecedence > precedence {
				// if operand1 is only a number, no need to wrap it in parentheses
				if _, err := strconv.Atoi(operand1); err != nil {
					operand1 = fmt.Sprintf("(%s)", operand1)
				}
				if _, err := strconv.Atoi(operand2); err != nil {
					operand2 = fmt.Sprintf("(%s)", operand2)
				}
			}
			prevPrecedence = precedence

			stack = append(stack, fmt.Sprintf("%s %s %s", operand1, operator, operand2))
		} else {
			stack = append(stack, token)
		}
	}

	return stack[0], nil
}
