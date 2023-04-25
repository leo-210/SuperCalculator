package parser

import "testing"

func TestTokenize(t *testing.T) {
	var result, err = Tokenize("1 \t2 3.4 5. .6 ... 7890 +-*/^()")

	if err != nil {
		t.Errorf("Got error: %s", err)
		return
	}

	var expected = []Token{
		{NUMBER, "1"},
		{NUMBER, "2"},
		{NUMBER, "3.4"},
		{NUMBER, "5"},
		{NUMBER, "0.6"},
		{NUMBER, "7890"},
		{PLUS, ""},
		{MINUS, ""},
		{MULTIPLY, ""},
		{DIVIDE, ""},
		{POWER, ""},
		{LPAREN, ""},
		{RPAREN, ""},
	}

	if len(result) != len(expected) {
		t.Errorf("Expected : %+v\nGot :      %+v", expected, result)
		return
	}

	for i, token := range result {
		if token.Type != expected[i].Type || token.Value != expected[i].Value {
			t.Errorf("Expected : %+v\nGot :      %+v", expected, result)
			return
		}
	}
}
