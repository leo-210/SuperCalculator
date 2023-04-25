package parser

import "testing"

func TestParse(t *testing.T) {
	var input = []Token{
		{Type: LPAREN},
		{Type: NUMBER, Value: "8"},
		{Type: PLUS},
		{Type: NUMBER, Value: "7"},
		{Type: RPAREN},
		{Type: POWER},
		{Type: NUMBER, Value: "2"},
		{Type: DIVIDE},
		{Type: LPAREN},
		{Type: NUMBER, Value: "3"},
		{Type: MINUS},
		{Type: NUMBER, Value: "5"},
		{Type: RPAREN},
	} // (8 + 7)^2 / (3 - 5)

	var result, err = Parse(input)

	if err != nil {
		t.Errorf("Got error: %s", err)
		return
	}

	var expected = Node{
		Type: DIV,
		Left: &Node{
			Type: POW,
			Left: &Node{
				Type: ADD,
				Left: &Node{
					Type:  VALUE,
					Value: "8",
				},
				Right: &Node{
					Type:  VALUE,
					Value: "7",
				},
			},
			Right: &Node{
				Type:  VALUE,
				Value: "2",
			},
		},
		Right: &Node{
			Type: SUB,
			Left: &Node{
				Type:  VALUE,
				Value: "3",
			},
			Right: &Node{
				Type:  VALUE,
				Value: "5",
			},
		},
	}

	if !compareAST(expected, result) {
		t.Errorf("Expected : %+v\nGot :      %+v", ASTToString(expected), ASTToString(result))
		return
	}
}

func compareAST(ast1 Node, ast2 Node) bool {
	if ast1.Type != ast2.Type {
		return false
	}

	if ast1.Type == VALUE {
		if ast1.Value != ast2.Value {
			return false
		}
	} else {
		if !compareAST(*ast1.Left, *ast2.Left) {
			return false
		}

		if !compareAST(*ast1.Right, *ast2.Right) {
			return false
		}
	}

	return true
}
