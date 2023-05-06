package derivator

import (
	"SuperCalculator/src/interpreter/calculator"
	"SuperCalculator/src/interpreter/simplifier"
	"SuperCalculator/src/parser"
)

func Derive(ast parser.Node, variables map[string]float64) (parser.Node, error) {
	var err error
	ast, err = calculator.Calculate(ast, variables)

	if err != nil {
		return ast, err
	}

	return simplifier.Simplify(deriveNode(ast), variables)
}

func deriveNode(node parser.Node) parser.Node {
	switch node.Type {
	case parser.ADD, parser.SUB:
		var left = deriveNode(*node.Left)
		var right = deriveNode(*node.Right)

		return parser.MakeOperationNode(node.Type, left, right)

	case parser.MUL:
		var left = deriveNode(*node.Left)
		var right = deriveNode(*node.Right)

		return parser.MakeOperationNode(
			parser.ADD,
			parser.MakeOperationNode(
				parser.MUL,
				left,
				*node.Right,
			),
			parser.MakeOperationNode(
				parser.MUL,
				*node.Left,
				right,
			),
		)

	case parser.DIV:
		var left = deriveNode(*node.Left)
		var right = deriveNode(*node.Right)

		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeOperationNode(
				parser.SUB,
				parser.MakeOperationNode(
					parser.MUL,
					left,
					*node.Right,
				),
				parser.MakeOperationNode(
					parser.MUL,
					*node.Left,
					right,
				),
			),
			parser.MakeOperationNode(
				parser.POW,
				*node.Right,
				parser.MakeValueNode("2"),
			),
		)

	case parser.POW:
		var left = deriveNode(*node.Left)
		var right = deriveNode(*node.Right)

		return parser.MakeOperationNode(
			parser.MUL,
			node,
			parser.MakeOperationNode(
				parser.ADD,
				parser.MakeOperationNode(
					parser.DIV,
					parser.MakeOperationNode(
						parser.MUL,
						left,
						*node.Right),
					*node.Left,
				),
				parser.MakeOperationNode(
					parser.MUL,
					right,
					parser.MakeFunctionNode("ln", *node.Left),
				),
			),
		)

	case parser.FUNCTION:
		var left = deriveNode(*node.Left)

		return parser.MakeOperationNode(
			parser.MUL,
			left,
			Derivatives[node.Value](*node.Left),
		)

	case parser.X:
		return parser.MakeValueNode("1")

	case parser.VALUE:
		return parser.MakeValueNode("0")

	// **Should** be unreachable
	default:
		return parser.Node{}

	}
}
