package calculator

import (
	"SuperCalculator/src/errors"
	"SuperCalculator/src/parser"
	"math"
	"strconv"
)

func Calculate(ast parser.Node) (float64, error) {
	switch ast.Type {
	case parser.ADD:
		var left, err = Calculate(*ast.Left)
		if err != nil {
			return 0, err
		}

		var right float64
		right, err = Calculate(*ast.Right)
		if err != nil {
			return 0, err
		}

		return left + right, nil

	case parser.SUB:
		var left, err = Calculate(*ast.Left)
		if err != nil {
			return 0, err
		}

		var right float64
		right, err = Calculate(*ast.Right)
		if err != nil {
			return 0, err
		}

		return left - right, nil

	case parser.MUL:
		var left, err = Calculate(*ast.Left)
		if err != nil {
			return 0, err
		}

		var right float64
		right, err = Calculate(*ast.Right)
		if err != nil {
			return 0, err
		}

		return left * right, nil

	case parser.DIV:
		var right, err = Calculate(*ast.Right)
		if err != nil {
			return 0, err
		}

		if right == 0 {
			return 0, errors.MathError{Message: "division by zero"}
		}

		var left float64
		left, err = Calculate(*ast.Left)
		if err != nil {
			return 0, err
		}

		return left / right, nil

	case parser.POW:
		var left, err = Calculate(*ast.Left)
		if err != nil {
			return 0, err
		}

		var right float64
		right, err = Calculate(*ast.Right)
		if err != nil {
			return 0, err
		}

		if left == 0 && right == 0 {
			return 0, errors.MathError{Message: "zero to the power of zero"}
		}

		return math.Pow(left, right), nil

	case parser.VALUE:
		var val, err = strconv.ParseFloat(ast.Value, 64)
		if err != nil {
			return 0, err
		}

		return val, nil

	default:
		return 0, errors.NotImplementedYetError{}
	}
}
