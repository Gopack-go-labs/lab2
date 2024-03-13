package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "empty input 2",
			input:    "  \n",
			expected: "",
		},
		{
			name:     "negative 1",
			input:    "4 -2 + 3 +",
			expected: "4 - 2 + 3",
		},
		{
			name:     "negative 2",
			input:    "4 -2 - 3 +",
			expected: "4 + 2 + 3",
		},
		{
			name:     "postfix 1",
			input:    "4 2 - 3 * 5 2 3 + ^ +",
			expected: "(4 - 2) * 3 + 5 ^ (2 + 3)",
		},
		{
			name:     "postfix 2 with spaces",
			input:    "    5     2 3    4            ^\t + ^    \n",
			expected: "5 ^ (2 + 3 ^ 4)",
		},
		{
			name:     "postfix 3",
			input:    "1 2 3 4 ^ ^ / 5 + 6 * 7 ^ 100 -",
			expected: "((1 / 2 ^ 3 ^ 4 + 5) * 6) ^ 7 - 100",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, err := PostfixToInfix(tt.input)
			if assert.Nil(t, err) {
				assert.Equal(t, tt.expected, res)
			}
		})
	}
}

func TestPostfixToInfixInvalid(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input string
	}{
		{
			name:  "Invalid expression",
			input: "5 4 ^ 3 6",
		},
		{
			name:  "Invalid symbol",
			input: "hello world +",
		},
		{
			name:  "Invalid symbol 2",
			input: " a \t b +",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := PostfixToInfix(tt.input)
			assert.NotNil(t, err)
		})
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
