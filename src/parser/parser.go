package parser

import (
	"SuperCalculator/src/defined_identifiers"
	"SuperCalculator/src/errors"
	"fmt"
	"strconv"
	"strings"
)

type NodeType int

const (
	ADD NodeType = iota
	SUB
	MUL
	DIV
	POW
	VALUE
	VARIABLE
	SET_VARIABLE
	FUNCTION
	X
	DERIVE
)

func (nodeType NodeType) String() string {
	return [...]string{
		"ADD",
		"SUB",
		"MUL",
		"DIV",
		"POW",
		"VALUE",
		"VARIABLE",
		"SET_VARIABLE",
		"FUNCTION",
		"X",
		"DERIVE",
	}[nodeType]
}

type Node struct {
	Left  *Node
	Right *Node
	Value string
	Type  NodeType
}

func Parse(tokenList []Token) (Node, error) {
	if len(tokenList) == 0 {
		return Node{}, nil
	}

	var nodeList, err = parseStatement(tokenList)

	if err != nil {
		return nodeList, err
	}

	return nodeList, nil
}

func parseStatement(tokenList []Token) (Node, error) {
	var token = tokenList[0]

	var node Node
	var err error
	var i int

	switch token.Type {
	case KEYWORD:
		switch token.Value {
		case "derive":
			var expr Node
			expr, i, err = parseExpression(tokenList, 1)

			node = Node{Type: DERIVE, Left: &expr, Value: token.Value}
		}
	case IDENTIFIER:
		if len(tokenList) > 2 && tokenList[1].Type == EQUAL {
			if (token.Value == "pi" || token.Value == "PI") && tokenList[2].Type == NUMBER && tokenList[2].Value == "5" {
				node = Node{Type: SET_VARIABLE, Left: &Node{Type: VALUE, Value: "5"}, Value: "pi"}
				return node, nil
			}

			var _, ok = defined_identifiers.Constants[token.Value]
			if !ok {
				_, ok = defined_identifiers.Functions[token.Value]
				if !ok {
					var expr Node
					expr, i, err = parseExpression(tokenList, 2)

					node = Node{Type: SET_VARIABLE, Left: &expr, Value: token.Value}
				} else {
					return Node{}, errors.AlreadyDefinedIdentifierError{IdentifierType: "function", Name: token.Value}
				}
			} else {
				return Node{}, errors.AlreadyDefinedIdentifierError{IdentifierType: "constant", Name: token.Value}
			}
		} else {
			node, i, err = parseExpression(tokenList, 0)
		}

	default:
		node, i, err = parseExpression(tokenList, 0)
	}

	if err != nil {
		return node, err
	}

	if i < len(tokenList) {
		return node, errors.SyntaxError{Message: "unexpected character", Position: i}
	}

	return node, nil
}

func parseExpression(tokenList []Token, index int) (Node, int, error) {
	var i = index
	var node Node

	if i >= len(tokenList) || tokenList[i].Type == RPAREN {
		return node, i, errors.ExpressionExpectedError{Position: i}
	}

	var token = tokenList[i]
	var err error

	if token.Type == PLUS || token.Type == MINUS || token.Type == LPAREN || token.Type == NUMBER || token.Type == IDENTIFIER || token.Type == T_X {
		node, i, err = parseTerm(tokenList, i)
		if err != nil {
			return node, i, err
		}
	} else if token.Type != LPAREN && token.Type != NUMBER {
		return node, i, errors.ExpressionExpectedError{Position: i}
	}

	for i < len(tokenList) {
		token = tokenList[i]

		switch token.Type {
		case PLUS:
			var firstNode = node
			var secondNode Node
			secondNode, i, err = parseTerm(tokenList, i+1)

			node = Node{Left: &firstNode, Right: &secondNode, Type: ADD}

			if err != nil {
				return node, i, err
			}
		case MINUS:
			var firstNode = node
			var secondNode Node
			secondNode, i, err = parseTerm(tokenList, i+1)

			node = Node{Left: &firstNode, Right: &secondNode, Type: SUB}

			if err != nil {
				return node, i, err
			}
		default:
			return node, i, nil

		}
	}

	return node, i + 1, nil
}

func parseTerm(tokenList []Token, index int) (Node, int, error) {
	var i = index

	if i >= len(tokenList) {
		return Node{}, i, errors.SyntaxError{Message: "syntax error", Position: i}
	}

	var token = tokenList[i]
	var node Node
	var err error

	if token.Type == PLUS || token.Type == MINUS || token.Type == LPAREN || token.Type == NUMBER || token.Type == IDENTIFIER || token.Type == T_X {
		node, i, err = parseFactor(tokenList, i)
		if err != nil {
			return node, i, err
		}
	} else {
		return node, i, errors.SyntaxError{Message: "syntax error", Position: i}
	}

	for i < len(tokenList) {
		token = tokenList[i]

		switch token.Type {

		case MULTIPLY:
			var firstNode = node
			var secondNode Node
			secondNode, i, err = parseFactor(tokenList, i+1)

			node = Node{Left: &firstNode, Right: &secondNode, Type: MUL}

			if err != nil {
				return node, i, err
			}
		case IDENTIFIER, LPAREN, T_X:
			var firstNode = node
			var secondNode Node
			secondNode, i, err = parseFactor(tokenList, i)

			node = Node{Left: &firstNode, Right: &secondNode, Type: MUL}

			if err != nil {
				return node, i, err
			}
		case DIVIDE:
			var firstNode = node
			var secondNode Node
			secondNode, i, err = parseFactor(tokenList, i+1)

			node = Node{Left: &firstNode, Right: &secondNode, Type: DIV}

			if err != nil {
				return node, i, err
			}
		default:
			return node, i, nil

		}
	}

	return node, i + 1, nil
}

func parseFactor(tokenList []Token, index int) (Node, int, error) {
	var i = index

	if i >= len(tokenList) {
		return Node{}, i, errors.SyntaxError{Message: "syntax error", Position: i}
	}

	var token = tokenList[i]
	var node Node
	var err error

	if token.Type == PLUS || token.Type == MINUS || token.Type == LPAREN || token.Type == NUMBER || token.Type == IDENTIFIER || token.Type == T_X {
		node, i, err = parseAtom(tokenList, i)
		if err != nil {
			return node, i, err
		}
	} else {
		return node, i, errors.SyntaxError{Message: "syntax error", Position: i}
	}

	for i < len(tokenList) {
		token = tokenList[i]

		switch token.Type {
		case POWER:
			var firstNode = node
			var secondNode Node
			secondNode, i, err = parseAtom(tokenList, i+1)

			if firstNode.Type == VALUE &&
				firstNode.Value == strconv.FormatFloat(defined_identifiers.Constants["e"], 'f', -1, 64) {
				node = MakeFunctionNode("exp", secondNode)
			} else {
				node = Node{Left: &firstNode, Right: &secondNode, Type: POW}
			}

			if err != nil {
				return node, i, err
			}

		default:
			return node, i, nil
		}
	}

	return node, i + 1, nil
}

func parseAtom(tokenList []Token, index int) (Node, int, error) {
	var i = index

	if i >= len(tokenList) {
		return Node{}, i, errors.SyntaxError{Message: "syntax error", Position: i}
	}

	var token = tokenList[i]
	var err error

	switch token.Type {
	case PLUS, MINUS:
		var firstNode Node

		if token.Type == PLUS {
			firstNode = Node{Type: VALUE, Value: "1"}
		} else {
			firstNode = Node{Type: VALUE, Value: "-1"}
		}
		var secondNode Node

		secondNode, i, err = parseAtom(tokenList, i+1)
		return Node{Type: MUL, Left: &firstNode, Right: &secondNode}, i, err

	case NUMBER:
		return Node{Type: VALUE, Value: token.Value}, i + 1, nil

	case LPAREN:
		var node Node
		node, i, err = parseExpression(tokenList, i+1)

		if i >= len(tokenList) || tokenList[i].Type != RPAREN {
			return node, i, errors.MissingClosingParenthesisError{Position: i}
		}
		i++

		return node, i, err
	case T_X:
		return Node{Type: X}, i + 1, nil
	case IDENTIFIER:
		var value, ok = defined_identifiers.Constants[token.Value]
		if ok {
			return Node{Type: VALUE, Value: fmt.Sprintf("%.14f", value)}, i + 1, nil
		}
		_, ok = defined_identifiers.Functions[token.Value]

		if ok {
			var body Node
			body, i, err = parseAtom(tokenList, i+1)
			return Node{Type: FUNCTION, Value: token.Value, Left: &body}, i, err
		}

		if strings.HasPrefix(token.Value, "sqrt") {
			var number, _ = strings.CutPrefix(token.Value, "sqrt")

			_, err = strconv.ParseInt(number, 10, 32)

			if err != nil {
				return Node{Type: VARIABLE, Value: token.Value}, i, nil
			}
			var base Node
			base, i, err = parseAtom(tokenList, i+1)

			return Node{
				Type: POW,
				Left: &base,
				Right: &Node{
					Type: DIV,
					Left: &Node{
						Type:  VALUE,
						Value: "1",
					},
					Right: &Node{
						Type:  VALUE,
						Value: number,
					},
				},
			}, i, err
		}

		return Node{Type: VARIABLE, Value: token.Value}, i + 1, nil

	default:
		return Node{}, i, errors.SyntaxError{Message: "syntax error", Position: i}
	}
}

func (a Node) Equals(b Node) bool {
	if a.Type != b.Type {
		return false
	}

	switch a.Type {
	case ADD, SUB, MUL, DIV, POW:
		return a.Left.Equals(*b.Left) && a.Right.Equals(*b.Right)
	case VALUE, VARIABLE:
		return a.Value == b.Value
	case FUNCTION:
		return a.Value == b.Value && a.Left.Equals(*b.Left)
	default:
		return true
	}
}

func (a Node) Decompose() (Node, Node) {
	return *a.Left, *a.Right
}

func (a Node) DecomposeFunc(fn func(node Node) Node) (Node, Node) {
	return fn(*a.Left), fn(*a.Right)
}

// Utility functions

func ASTToString(ast Node) string {
	switch ast.Type {
	case ADD:
		return fmt.Sprintf("(%s + %s)", ASTToString(*ast.Left), ASTToString(*ast.Right))
	case SUB:
		return fmt.Sprintf("(%s - %s)", ASTToString(*ast.Left), ASTToString(*ast.Right))
	case MUL:
		if ast.Left.Type == VALUE && ast.Left.Value == "-1" {
			return fmt.Sprintf("-%s", ASTToString(*ast.Right))
		}
		if ast.Right.Type == VALUE && ast.Right.Value == "-1" {
			return fmt.Sprintf("(-%s)", ASTToString(*ast.Left))
		}

		return fmt.Sprintf("(%s * %s)", ASTToString(*ast.Left), ASTToString(*ast.Right))
	case DIV:
		return fmt.Sprintf("(%s / %s)", ASTToString(*ast.Left), ASTToString(*ast.Right))
	case POW:
		return fmt.Sprintf("(%s ^ %s)", ASTToString(*ast.Left), ASTToString(*ast.Right))
	case VALUE:
		return ast.Value
	case VARIABLE:
		return ast.Value
	case FUNCTION:
		var resultS, ok = strings.CutPrefix(ASTToString(*ast.Left), "(")
		if ok {
			resultS = strings.TrimSuffix(resultS, ")")
		}

		return fmt.Sprintf("%s(%s)", ast.Value, resultS)
	case SET_VARIABLE:
		return fmt.Sprintf("%s = %s", ast.Value, ASTToString(*ast.Left))
	case X:
		return "x"
	case DERIVE:
		return fmt.Sprintf("derive %s", ASTToString(*ast.Left))
	default:
		return ""
	}
}

func MakeValueNode(value string) Node {
	return Node{
		Type:  VALUE,
		Value: value,
	}
}

func MakeOperationNode(op NodeType, left Node, right Node) Node {
	return Node{
		Type:  op,
		Left:  &left,
		Right: &right,
	}
}

func MakeFunctionNode(function string, body Node) Node {
	return Node{
		Type:  FUNCTION,
		Value: function,
		Left:  &body,
	}
}
