package main

import (
	"flag"
	"fmt"

	"github.com/Gopack-go-labs/lab2"
)

var (
	expressionFlag = flag.String("e", "", "Postfix expression")
	fileFlag       = flag.String("f", "", "File containing postfix expression")
	outputFlag     = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, _ := lab2.PostfixToInfix("+ 2 2")
	fmt.Println(res)
}
