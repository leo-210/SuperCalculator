package simplifier

import (
	"SuperCalculator/src/interpreter/calculator"
	"SuperCalculator/src/parser"
)

func Simplify(ast parser.Node, variables map[string]float64) (parser.Node, error) {
	var node = ast
	for i := 0; i < 10; i++ {
		var newNode = regroup(node)
		if newNode.Equals(node) {
			break
		}
		node, _ = calculator.Calculate(newNode, variables)
	}
	return calculator.Calculate(node, variables)
}
