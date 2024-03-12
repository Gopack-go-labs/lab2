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
			name:     "postfix 2",
			input:    "5 2 3 4 ^ + ^",
			expected: "5 ^ (2 + 3 ^ 4)",
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

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
