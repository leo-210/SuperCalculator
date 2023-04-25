package interpreter

import (
	"SuperCalculator/src/interpreter/calculator"
	"SuperCalculator/src/parser"
	"fmt"
	"strings"
)

func Interpret(ast parser.Node) (string, error) {
	var result, err = calculator.Calculate(ast)
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", result), "0"), "."), err
}
