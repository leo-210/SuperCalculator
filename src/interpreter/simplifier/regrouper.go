package simplifier

import (
	"SuperCalculator/src/parser"
	"math"
	"strconv"
)

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
		} else if right.Type == parser.VALUE {
			var f, _ = strconv.ParseFloat(right.Value, 64)
			if f == math.Round(f) {
				if f == 0 {
					node = parser.MakeValueNode("1")
				} else if f == 1 {
					node = left
				} else if f > 1 {
					node = left

					// TODO : implement (x + y)^n and (x - y)^n formulas

					for i := 1; i < int(f); i++ {
						node = parser.MakeOperationNode(
							parser.MUL,
							left,
							node,
						)
					}
				}
			}
		}

	case parser.MUL:
		ast = distribute(ast)
		if ast.Type != parser.MUL {
			return regroup(ast)
		}

		left, right = ast.DecomposeFunc(distribute)
		left, right = regroup(left), regroup(right)

		if left.Type == parser.X && right.Type == parser.VALUE {
			return parser.MakeOperationNode(
				parser.MUL,
				right,
				left,
			)
		}

		if left.Type == parser.DIV {
			return regroup(parser.MakeOperationNode(
				parser.DIV,
				parser.MakeOperationNode(
					parser.MUL,
					*left.Left,
					right,
				),
				*left.Right,
			))
		}
		if right.Type == parser.DIV {
			return regroup(parser.MakeOperationNode(
				parser.DIV,
				parser.MakeOperationNode(
					parser.MUL,
					*right.Left,
					left,
				),
				*right.Right,
			))
		}

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
			newRight = *newRight.Left
		}

		node = parser.MakeOperationNode(
			parser.MUL,
			newLeft,
			newRight,
		)

		if newRight.Equals(parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			parser.MakeValueNode("1"),
		)) || newRight.Equals(parser.MakeValueNode("1")) {
			node = newLeft
		}

	case parser.DIV:
		var newLeft, newRight = parser.Node{Type: parser.DIV}, parser.Node{Type: parser.DIV}

		if left.Type == parser.MUL {
			if left.Left.Equals(right) {
				return *left.Right
			}
			if left.Right.Equals(right) {
				return *left.Left
			}

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
				parser.DIV,
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
			newRight = *newRight.Left
		}

		if left.Equals(right) {
			node = parser.MakeValueNode("1")
		} else {
			node = parser.MakeOperationNode(
				parser.MUL,
				newLeft,
				newRight,
			)
		}

		if newRight.Equals(parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			parser.MakeValueNode("1"),
		)) || newRight.Equals(parser.MakeValueNode("1")) {
			node = newLeft
		}

		if node.Right.Type == parser.DIV && node.Right.Left.Equals(*node.Right.Right) {
			node = *node.Left
		}
	case parser.ADD:
		node = parser.MakeOperationNode(
			node.Type,
			left,
			right,
		)
	case parser.SUB:
		node = parser.MakeOperationNode(
			node.Type,
			left,
			right,
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
