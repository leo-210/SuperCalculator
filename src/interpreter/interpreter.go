package interpreter

import (
	"SuperCalculator/src/defined_identifiers"
	"SuperCalculator/src/interpreter/calculator"
	"SuperCalculator/src/parser"
	"fmt"
	"math"
	"strings"
)

func Interpret(ast parser.Node, variables map[string]float64, mode int) (string, map[string]float64, error) {
	if ast == (parser.Node{}) {
		return "", variables, nil
	}

	switch ast.Type {
	case parser.SET_VARIABLE:
		if ast.Value == "pi" {
			defined_identifiers.Constants["pi"] = 5
			defined_identifiers.Constants["PI"] = 5
			return "engineer-mode", variables, nil
		}

		var err error
		variables[ast.Value], err = calculator.Calculate(*ast.Left, variables)

		return "set " + ast.Value, variables, err
	}

	var result, err = calculator.Calculate(ast, variables)

	if mode == 1 {
		result = math.Round(result)
	}

	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", result), "0"), "."), variables, err
}
