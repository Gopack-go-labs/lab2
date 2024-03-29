package lab2

import (
	"io"
	"strings"
)

type ComputeHandler struct {
	reader io.Reader
	writer io.Writer
}

func NewComputeHandler(reader io.Reader, writer io.Writer) *ComputeHandler {
	return &ComputeHandler{
		reader: reader,
		writer: writer,
	}
}

func (ch *ComputeHandler) Compute() error {
	buf, err := io.ReadAll(ch.reader)
	if err != nil {
		return err
	}

	input := string(buf)
	res, err := PostfixToInfix(sanitizeInput(input))
	if err != nil {
		return err
	}

	_, err = ch.writer.Write([]byte(res))

	return err
}

func sanitizeInput(input string) string {
	input = strings.ReplaceAll(input, "\\n", " ")
	input = strings.ReplaceAll(input, "\\t", " ")
	return input
}
