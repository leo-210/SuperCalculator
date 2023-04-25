package main

import (
	"SuperCalculator/src/interpreter"
	"SuperCalculator/src/parser"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {
	var cyan = color.New(color.FgCyan).SprintFunc()
	var green = color.New(color.FgGreen).SprintFunc()

	fmt.Printf("Welcome to the %s version %s ! Type something for the calculator to do.\n"+
		"For help, type %s. To quit the application, type %s.\n\n",
		green("SuperCalculator"),
		green("0.1.0"),
		cyan("help"),
		cyan("quit"),
	)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s > ", green("SuperCalculator"))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		if text == "exit" || text == "quit" || text == "q" || text == "Q" {
			os.Exit(0)
		}

		if text == "help" || text == "h" || text == "H" {
			fmt.Printf("The %s menu is not implemented yet.\n", cyan("help"))
			continue
		}

		tokens, err := parser.Tokenize(text)
		if err != nil {
			color.Red("error: %s", err)
			continue
		}

		var ast parser.Node
		ast, err = parser.Parse(tokens)
		if err != nil {
			color.Red("error: %s", err)
			continue
		}

		var result string
		result, err = interpreter.Interpret(ast)
		if err != nil {
			color.Red("error: %s", err)
			continue
		}

		fmt.Printf("%s = %s\n", text, cyan(result))
	}
}
