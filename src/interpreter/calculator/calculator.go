package calculator

import (
	"SuperCalculator/src/defined_identifiers"
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

	case parser.FUNCTION:
		var function, _ = defined_identifiers.Functions[ast.Value]
		var body, err = Calculate(*ast.Left)

		if err != nil {
			return 0, err
		}

		if ast.Value == "sqrt" {
			if body < 0 {
				return 0, errors.MathError{Message: "square root of a non-zero negative number"}
			}
		} else if ast.Value == "ln" || ast.Value == "log" {
			if body <= 0 {
				return 0, errors.MathError{Message: "natural logarithm of a negative number"}
			}
		} else if ast.Value == "asin" || ast.Value == "acos" || ast.Value == "atanh" {
			if -1.0 > body || body > 1 {
				return 0, errors.MathError{Message: "asin/acos/atanh of a number not between -1 and 1"}
			}
		} else if ast.Value == "acosh" {
			if body < 1 {
				return 0, errors.MathError{Message: "acosh of a number smaller than 1"}
			}
		}

		return function(body), nil
	default:
		return 0, errors.NotImplementedYetError{}
	}
}
