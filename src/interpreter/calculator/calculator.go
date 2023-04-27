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

func Calculate(ast parser.Node, variables map[string]float64, mode int) (string, error) {
	var result float64

	switch ast.Type {
	case parser.ADD:
		if ast.Left.Type == parser.X && ast.Right.Type == parser.X {
			return parser.ASTToString(ast), nil
		}
		if ast.Left.Type == parser.X {
			var right, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(x + %s)", right), nil
		}
		if ast.Right.Type == parser.X {
			var left, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(%s + x)", left), nil
		}

		var left, err = Calculate(*ast.Left, variables, mode)
		if err != nil {
			return "", err
		}

		var right string
		right, err = Calculate(*ast.Right, variables, mode)
		if err != nil {
			return "", err
		}

		if strings.Contains(left, "x") || strings.Contains(right, "x") {
			return parser.ASTToString(ast), nil
		}

		var leftF, _ = strconv.ParseFloat(left, 64)
		var rightF, _ = strconv.ParseFloat(right, 64)

		result = leftF + rightF

	case parser.SUB:
		if ast.Left.Type == parser.X && ast.Right.Type == parser.X {
			return parser.ASTToString(ast), nil
		}
		if ast.Left.Type == parser.X {
			var right, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(x - %s)", right), nil
		}
		if ast.Right.Type == parser.X {
			var left, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(%s - x)", left), nil
		}

		var left, err = Calculate(*ast.Left, variables, mode)
		if err != nil {
			return "", err
		}

		var right string
		right, err = Calculate(*ast.Right, variables, mode)
		if err != nil {
			return "", err
		}

		if strings.Contains(left, "x") || strings.Contains(right, "x") {
			return parser.ASTToString(ast), nil
		}

		var leftF, _ = strconv.ParseFloat(left, 64)
		var rightF, _ = strconv.ParseFloat(right, 64)

		result = leftF - rightF

	case parser.MUL:
		if ast.Left.Type == parser.X && ast.Right.Type == parser.X {
			return parser.ASTToString(ast), nil
		}
		if ast.Left.Type == parser.X {
			var right, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(x * %s)", right), nil
		}
		if ast.Right.Type == parser.X {
			var left, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(%s * x)", left), nil
		}

		var left, err = Calculate(*ast.Left, variables, mode)
		if err != nil {
			return "", err
		}

		var right string
		right, err = Calculate(*ast.Right, variables, mode)
		if err != nil {
			return "", err
		}

		if strings.Contains(left, "x") || strings.Contains(right, "x") {
			return parser.ASTToString(ast), nil
		}

		var leftF, _ = strconv.ParseFloat(left, 64)
		var rightF, _ = strconv.ParseFloat(right, 64)

		result = leftF * rightF

	case parser.DIV:
		if ast.Left.Type == parser.X && ast.Right.Type == parser.X {
			return parser.ASTToString(ast), nil
		}
		if ast.Left.Type == parser.X {
			var right, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(x / %s)", right), nil
		}
		if ast.Right.Type == parser.X {
			var left, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(%s / x)", left), nil
		}

		var right, err = Calculate(*ast.Right, variables, mode)
		if err != nil {
			return "", err
		}

		if strings.Contains(right, "x") {
			return parser.ASTToString(ast), nil
		}

		var rightF, _ = strconv.ParseFloat(right, 64)

		if rightF == 0 {
			return "", errors.MathError{Message: "division by zero"}
		}

		var left string
		left, err = Calculate(*ast.Left, variables, mode)
		if err != nil {
			return "", err
		}

		if strings.Contains(left, "x") {
			return parser.ASTToString(ast), nil
		}

		var leftF, _ = strconv.ParseFloat(left, 64)

		result = leftF / rightF

	case parser.POW:
		if ast.Left.Type == parser.X && ast.Right.Type == parser.X {
			return parser.ASTToString(ast), nil
		}
		if ast.Left.Type == parser.X {
			var right, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(x ^ %s)", right), nil
		}
		if ast.Right.Type == parser.X {
			var left, err = Calculate(*ast.Right, variables, mode)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("(%s ^ x)", left), nil
		}

		var left, err = Calculate(*ast.Left, variables, mode)
		if err != nil {
			return "", err
		}

		var right string
		right, err = Calculate(*ast.Right, variables, mode)
		if err != nil {
			return "", err
		}

		if strings.Contains(left, "x") || strings.Contains(right, "x") {
			return parser.ASTToString(ast), nil
		}

		var leftF, _ = strconv.ParseFloat(left, 64)
		var rightF, _ = strconv.ParseFloat(right, 64)

		if leftF == 0 && rightF == 0 {
			return "", errors.MathError{Message: "zero to the power of zero"}
		}

		result = math.Pow(leftF, rightF)

	case parser.VALUE:
		var val, err = strconv.ParseFloat(ast.Value, 64)
		if err != nil {
			return "", err
		}

		result = val

	case parser.FUNCTION:
		if ast.Left.Type == parser.X {
			return parser.ASTToString(ast), nil
		}

		var function, _ = defined_identifiers.Functions[ast.Value]
		var bodySTR, err = Calculate(*ast.Left, variables, mode)

		if strings.Contains(bodySTR, "x") {
			return parser.ASTToString(ast), nil
		}

		var body, _ = strconv.ParseFloat(bodySTR, 64)

		if err != nil {
			return "", err
		}

		if ast.Value == "sqrt" {
			if body < 0 {
				return "", errors.MathError{Message: "square root of a non-zero negative number"}
			}
		} else if ast.Value == "ln" || ast.Value == "log" {
			if body <= 0 {
				return "", errors.MathError{Message: "natural logarithm of a negative number"}
			}
		} else if ast.Value == "asin" || ast.Value == "acos" || ast.Value == "atanh" {
			if -1.0 > body || body > 1 {
				return "", errors.MathError{Message: "asin/acos/atanh of a number not between -1 and 1"}
			}
		} else if ast.Value == "acosh" {
			if body < 1 {
				return "", errors.MathError{Message: "acosh of a number smaller than 1"}
			}
		}

		result = function(body)

	case parser.VARIABLE:
		var value, ok = variables[ast.Value]
		if !ok {
			return "", errors.UndefinedVariableError{Name: ast.Value}
		}

		result = value
	case parser.X:
		return "x", nil
	default:
		return "", errors.NotImplementedYetError{}
	}
	if mode == 1 {
		return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", math.Round(result)), "0"), "."), nil
	}
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", result), "0"), "."), nil
}
