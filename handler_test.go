package lab2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestComputeHandler(t *testing.T) {
	input := "+ 2 2"
	computed, _ := PostfixToInfix(input)
	buffWriter := bytes.NewBuffer([]byte{})

	handler := NewComputeHandler(strings.NewReader(input), buffWriter)
	_ = handler.Compute()

	assert.Equal(t, computed, buffWriter.String())
}
func TestComputeHandlerError(t *testing.T) {
	input := "+ 2 + 2"
	buffWriter := bytes.NewBuffer([]byte{})

	handler := NewComputeHandler(strings.NewReader(input), buffWriter)
	err := handler.Compute()

	assert.NotNil(t, err)
}
