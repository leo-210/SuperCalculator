package calculator

import (
	"SuperCalculator/src/defined_identifiers"
	"SuperCalculator/src/errors"
	"SuperCalculator/src/parser"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Calculate(ast parser.Node, variables map[string]float64) (parser.Node, error) {
	switch ast.Type {
	case parser.ADD, parser.SUB, parser.MUL, parser.DIV, parser.POW:
		if ast.Left.Type == parser.X && ast.Right.Type == parser.X {
			if ast.Type == parser.SUB {
				return parser.MakeValueNode("0"), nil
			}

			return ast, nil
		}

		var left, err = Calculate(*ast.Left, variables)
		if err != nil {
			return parser.Node{}, err
		}

		var right parser.Node
		right, err = Calculate(*ast.Right, variables)
		if err != nil {
			return parser.Node{}, err
		}

		if left.Type != parser.VALUE && right.Type != parser.VALUE {
			return parser.MakeOperationNode(
				ast.Type,
				left,
				right,
			), nil
		}
		if right.Type == parser.VALUE && left.Type != parser.VALUE {
			if right.Value == "0" {
				if ast.Type == parser.ADD || ast.Type == parser.SUB {
					return left, nil
				}
				if ast.Type == parser.MUL {
					return parser.MakeValueNode("0"), nil
				}
				if ast.Type == parser.DIV {
					return parser.Node{}, errors.MathError{Message: "division by zero"}
				}
			} else if right.Value == "1" {
				if ast.Type == parser.MUL || ast.Type == parser.DIV {
					return left, nil
				}
			}

			return parser.Node{
				Type:  ast.Type,
				Left:  &left,
				Right: &right,
			}, nil
		}
		if left.Type == parser.VALUE && right.Type != parser.VALUE {
			if left.Value == "0" {
				if ast.Type == parser.ADD {
					return right, nil
				}
				if ast.Type == parser.SUB {
					return parser.MakeOperationNode(
						parser.MUL,
						parser.MakeValueNode("-1"),
						right,
					), nil
				}
				if ast.Type == parser.MUL || ast.Type == parser.DIV {
					return parser.MakeValueNode("0"), nil
				}
			} else if left.Value == "1" {
				if ast.Type == parser.MUL {
					return right, nil
				}
			}

			return parser.Node{
				Type:  ast.Type,
				Left:  &left,
				Right: &right,
			}, nil
		}

		return processValues(ast.Type, left, right)

	case parser.FUNCTION:
		if ast.Left.Type == parser.X {
			return ast, nil
		}

		var function, _ = defined_identifiers.Functions[ast.Value]
		var body, err = Calculate(*ast.Left, variables)

		if err != nil {
			return parser.Node{}, err
		}

		if body.Type != parser.VALUE {
			return parser.MakeFunctionNode(ast.Value, body), nil
		}

		var bodyF, _ = strconv.ParseFloat(body.Value, 64)

		if err != nil {
			return parser.Node{}, err
		}

		if ast.Value == "sqrt" {
			if bodyF < 0 {
				return parser.Node{}, errors.MathError{Message: "square root of a non-zero negative number"}
			}
		} else if ast.Value == "ln" || ast.Value == "log" {
			if bodyF <= 0 {
				return parser.Node{}, errors.MathError{Message: "natural logarithm of a negative number"}
			}
		} else if ast.Value == "asin" || ast.Value == "acos" || ast.Value == "atanh" {
			if -1.0 > bodyF || bodyF > 1 {
				return parser.Node{}, errors.MathError{Message: "asin/acos/atanh of a number not between -1 and 1"}
			}
		} else if ast.Value == "acosh" {
			if bodyF < 1 {
				return parser.Node{}, errors.MathError{Message: "acosh of a number smaller than 1"}
			}
		}

		var resultS = strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", function(bodyF)), "0"), ".")
		return parser.MakeValueNode(resultS), nil

	case parser.VALUE:
		var _, err = strconv.ParseFloat(ast.Value, 64)
		if err != nil {
			return parser.Node{}, err
		}

		return ast, nil

	case parser.VARIABLE:
		var value, ok = variables[ast.Value]
		if !ok {
			return parser.Node{}, errors.UndefinedVariableError{Name: ast.Value}
		}

		var valueS = strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", value), "0"), ".")
		return parser.MakeValueNode(valueS), nil

	case parser.X:
		return ast, nil
	default:
		return parser.Node{}, errors.NotImplementedYetError{}
	}
}

func processValues(op parser.NodeType, left parser.Node, right parser.Node) (parser.Node, error) {
	var leftF, err = strconv.ParseFloat(left.Value, 64)
	if err != nil {
		return parser.Node{}, err
	}

	var rightF float64
	rightF, err = strconv.ParseFloat(right.Value, 64)
	if err != nil {
		return parser.Node{}, err
	}

	var result float64

	switch op {
	case parser.ADD:
		result = leftF + rightF
	case parser.SUB:
		result = leftF - rightF
	case parser.MUL:
		result = leftF * rightF
	case parser.DIV:
		if rightF == 0 {
			return parser.Node{}, errors.MathError{Message: "division by zero"}
		}
		result = leftF / rightF
	case parser.POW:
		if leftF == 0 && rightF == 0 {
			return parser.Node{}, errors.MathError{Message: "zero to the power of zero"}
		}
		result = math.Pow(leftF, rightF)
	}

	var resultS = strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", result), "0"), ".")
	return parser.MakeValueNode(resultS), nil
}
