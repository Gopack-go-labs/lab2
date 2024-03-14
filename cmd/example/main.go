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

	var input io.Reader
	var output io.Writer

	if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			os.Stderr.WriteString("Error opening file\n")
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else if *expressionFlag != "" {
		input = strings.NewReader(*expressionFlag)
	} else {
		os.Stderr.WriteString("No expression or file provided\n")
		os.Exit(1)
	}

	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			os.Stderr.WriteString("Error creating output file\n")
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := lab2.NewComputeHandler(input, output)
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
