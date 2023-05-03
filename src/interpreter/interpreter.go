package interpreter

import (
	"SuperCalculator/src/defined_identifiers"
	"SuperCalculator/src/errors"
	"SuperCalculator/src/interpreter/calculator"
	"SuperCalculator/src/interpreter/derivator"
	"SuperCalculator/src/parser"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ResultType int

const (
	SET_VARIABLE ResultType = iota
	VALUE
	EXPRESSION
	ENGINEER_MODE
	EMPTY
)

type Result struct {
	Type       ResultType
	Value      string
	ValueFloat float64
	VarName    string
}

func Interpret(ast parser.Node, variables map[string]float64, mode int) (Result, map[string]float64, error) {
	if ast == (parser.Node{}) {
		return Result{Type: EMPTY}, variables, nil
	}

	var err error
	var result parser.Node

	switch ast.Type {
	case parser.SET_VARIABLE:
		if ast.Value == "pi" {
			defined_identifiers.Constants["pi"] = 5
			defined_identifiers.Constants["PI"] = 5
			return Result{Type: ENGINEER_MODE}, variables, nil
		}

		var value, err = calculator.Calculate(*ast.Left, variables)

		if value.Type != parser.VALUE {
			return Result{}, variables, errors.SyntaxError{
				Message:  "can't assign an expression to a variable",
				Position: 0,
			}
		}

		variables[ast.Value], _ = strconv.ParseFloat(value.Value, 64)
		return Result{
			Type:       SET_VARIABLE,
			Value:      value.Value,
			ValueFloat: variables[ast.Value],
			VarName:    ast.Value,
		}, variables, err

	case parser.DERIVE:
		result, err = derivator.Derive(*ast.Left, variables)
		if err != nil {
			return Result{}, variables, err
		}

	default:
		result, err = calculator.Calculate(ast, variables)
		if err != nil {
			return Result{}, variables, err
		}

		var resultS, ok = strings.CutPrefix(parser.ASTToString(result), "(")
		if ok {
			resultS = strings.TrimSuffix(resultS, ")")
		}

		return Result{
			Type:  EXPRESSION,
			Value: resultS,
		}, variables, err
	}

	if result.Type == parser.VALUE {
		var resultF float64
		resultF, _ = strconv.ParseFloat(result.Value, 64)

		if mode == 1 {
			resultF = math.Round(resultF)
		}

		result.Value = strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.14f", resultF), "0"), ".")

		return Result{
			Type:       VALUE,
			Value:      result.Value,
			ValueFloat: resultF,
		}, variables, nil
	}

	var resultS, ok = strings.CutPrefix(parser.ASTToString(result), "(")
	if ok {
		resultS = strings.TrimSuffix(resultS, ")")
	}

	return Result{
		Type:  EXPRESSION,
		Value: resultS,
	}, variables, err
}
