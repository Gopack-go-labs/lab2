package main

import (
	"flag"
	"io"
	"os"

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

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	handler := lab2.NewComputeHandler(input, output)
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
