package simplifier

import (
	"SuperCalculator/src/interpreter/calculator"
	"SuperCalculator/src/parser"
)

func Simplify(ast parser.Node, variables map[string]float64) (parser.Node, error) {
	return calculator.Calculate(regroup(ast), variables)
}
