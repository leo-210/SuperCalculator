package simplifier

import "SuperCalculator/src/parser"

func regroup(ast parser.Node) parser.Node {
	var node = ast

	// Easier to write than a big if
	switch ast.Type {
	case parser.ADD, parser.SUB, parser.MUL, parser.DIV, parser.POW:
		// Do nothing

	default:
		return ast
	}

	var left, right = ast.DecomposeFunc(regroup)

	switch ast.Type {
	case parser.POW:
		if left.Type == parser.POW {
			node = parser.MakeOperationNode(
				parser.POW,
				*left.Left,
				regroup(parser.MakeOperationNode(
					parser.MUL,
					*left.Right,
					right,
				)),
			) // ((3 * x) ^ 2) ^ 2 => (3^2 * x^2)^2 =>
		} else if left.Type == parser.MUL {
			node = regroup(parser.MakeOperationNode(
				parser.MUL,
				regroup(parser.MakeOperationNode(
					parser.POW,
					*left.Left,
					right,
				)),
				regroup(parser.MakeOperationNode(
					parser.POW,
					*left.Right,
					right,
				)),
			))
		} else if left.Type == parser.DIV {
			node = regroup(parser.MakeOperationNode(
				parser.DIV,
				regroup(parser.MakeOperationNode(
					parser.POW,
					*left.Left,
					right,
				)),
				regroup(parser.MakeOperationNode(
					parser.POW,
					*left.Right,
					right,
				)),
			))
		}

	case parser.MUL:
		ast = distribute(ast)
		if ast.Type != parser.MUL {
			return regroup(ast)
		}

		left, right = ast.DecomposeFunc(distribute)

		var newLeft, newRight = parser.Node{Type: parser.MUL}, parser.Node{Type: parser.MUL}

		if left.Type == parser.MUL {
			if left.Right.Type == parser.VALUE {
				newLeft.Left = left.Right
				newRight.Left = left.Left
			} else {
				newLeft.Left = left.Left
				newRight.Left = left.Right
			}
		} else {
			newLeft.Left = &left
			newRight = parser.MakeOperationNode(
				parser.MUL,
				parser.MakeValueNode("1"),
				parser.MakeValueNode("1"),
			)
		}

		if right.Type == parser.MUL {
			if right.Right.Type == parser.VALUE {
				newLeft.Right = right.Right
				newRight.Right = right.Left
			} else {
				newLeft.Right = right.Left
				newRight.Right = right.Right
			}
		} else {
			newLeft.Right = &right
			newRight = parser.MakeOperationNode(
				parser.MUL,
				*newRight.Left,
				parser.MakeValueNode("1"),
			)
		}

		node = parser.MakeOperationNode(
			parser.MUL,
			newLeft,
			newRight,
		)
	}
	return node
}

func distribute(node parser.Node) parser.Node {
	// Easier to write than a big if
	if node.Type != parser.MUL {
		return node
	}

	if (node.Left.Type == parser.ADD || node.Left.Type == parser.SUB) && (node.Right.Type == parser.ADD || node.Right.Type == parser.SUB) {
		return parser.MakeOperationNode(
			node.Left.Type,
			distribute(parser.MakeOperationNode(
				parser.MUL,
				*node.Left.Left,
				*node.Right,
			)),
			distribute(parser.MakeOperationNode(
				parser.MUL,
				*node.Left.Right,
				*node.Right,
			)),
		)
	}
	if node.Right.Type == parser.ADD || node.Right.Type == parser.SUB {
		return parser.MakeOperationNode(
			node.Right.Type,
			parser.MakeOperationNode(
				parser.MUL,
				*node.Left,
				*node.Right.Left,
			),
			parser.MakeOperationNode(
				parser.MUL,
				*node.Left,
				*node.Right.Right,
			),
		)
	}
	if node.Left.Type == parser.ADD || node.Left.Type == parser.SUB {
		return parser.MakeOperationNode(
			node.Left.Type,
			parser.MakeOperationNode(
				parser.MUL,
				*node.Right,
				*node.Left.Left,
			),
			parser.MakeOperationNode(
				parser.MUL,
				*node.Right,
				*node.Left.Right,
			),
		)
	}

	return node
}
