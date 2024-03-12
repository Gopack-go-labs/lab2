package lab2

import "io"

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

	res, err := PostfixToInfix(string(buf))
	if err != nil {
		return err
	}

	_, err = ch.writer.Write([]byte(res))

	return err
}
