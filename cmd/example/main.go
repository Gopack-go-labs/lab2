package main

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/Gopack-go-labs/lab2"
)

var (
	expressionFlag = flag.String("e", "", "Postfix expression")
	fileFlag       = flag.String("f", "", "File containing postfix expression")
	outputFlag     = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	input := getInputReader()
	output := getOutputWriter()

	handler := lab2.NewComputeHandler(input, output)
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func getInputReader() io.Reader {
	var input io.Reader
	if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			os.Stderr.WriteString("Error opening file\n")
			os.Exit(1)
		}
		input = file
	} else if *expressionFlag != "" {
		input = strings.NewReader(*expressionFlag)
	} else {
		os.Stderr.WriteString("No expression or file provided\n")
		os.Exit(1)
	}
	return input
}

func getOutputWriter() io.Writer {
	var output io.Writer
	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			os.Stderr.WriteString("Error creating output file\n")
			os.Exit(1)
		}
		output = file
	} else {
		output = os.Stdout
	}
	return output
}
