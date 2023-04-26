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

	var variables = make(map[string]float64)
	var mode = 0
	var previousAnswer = "0"
	var previousPrompt = ""

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s > ", green("SuperCalculator"))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if len(text) == 0 && len(previousPrompt) != 0 {
			text = previousPrompt
		}

		if text == "exit" || text == "quit" || text == "q" || text == "Q" {
			os.Exit(0)
		}

		if text == "help" || text == "h" || text == "H" {
			fmt.Printf("The %s menu is not implemented yet.\n", cyan("help"))
			continue
		}

		text = strings.ReplaceAll(text, "ans", previousAnswer)

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
		result, variables, err = interpreter.Interpret(ast, variables, mode)
		if err != nil {
			color.Red("error: %s", err)
			continue
		}

		previousPrompt = text

		if result == "engineer-mode" {
			color.Red("Engineer mode activated. All results will be rounded.")
			mode = 1
		} else if strings.HasPrefix(result, "set ") {
			var variable, _ = strings.CutPrefix(result, "set ")
			var value = strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", variables[variable]), "0"), ".")
			fmt.Printf("%s = %s\n", variable, cyan(value))
		} else if result != "" {
			fmt.Printf("%s = %s\n", text, cyan(result))
			previousAnswer = result
		}
	}
}
