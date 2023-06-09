package parser

import (
	"SuperCalculator/src/defined_identifiers"
	"SuperCalculator/src/errors"
	"strings"
	"unicode"
)

type TokenType int

const (
	PLUS TokenType = iota
	MINUS
	MULTIPLY
	DIVIDE
	POWER
	LPAREN
	RPAREN
	NUMBER
	IDENTIFIER
	EQUAL
	KEYWORD
	T_X
)

func (tokenType TokenType) String() string {
	return [...]string{
		"PLUS",
		"MINUS",
		"MULTIPLY",
		"DIVIDE",
		"POWER",
		"LPAREN",
		"RPAREN",
		"NUMBER",
		"IDENTIFIER",
		"EQUAL",
		"KEYWORD",
		"T_X",
	}[tokenType]
}

type Token struct {
	Type  TokenType
	Value string
}

const decimalSeparator = "."

func Tokenize(textInput string) ([]Token, error) {
	var i = 0
	var tokenList []Token

	for i < len(textInput) {
		var char = textInput[i]

		switch {
		case char == '+':
			tokenList = append(tokenList, Token{PLUS, ""})
		case char == '-':
			tokenList = append(tokenList, Token{MINUS, ""})
		case char == '*':
			tokenList = append(tokenList, Token{MULTIPLY, ""})
		case char == '/':
			tokenList = append(tokenList, Token{DIVIDE, ""})
		case char == '^':
			tokenList = append(tokenList, Token{POWER, ""})
		case char == '(':
			tokenList = append(tokenList, Token{LPAREN, ""})
		case char == ')':
			tokenList = append(tokenList, Token{RPAREN, ""})
		case char == '=':
			tokenList = append(tokenList, Token{EQUAL, ""})
		case char == 'x':
			tokenList = append(tokenList, Token{T_X, ""})

		case unicode.IsSpace(rune(char)):
			// Do nothing

		case unicode.IsDigit(rune(char)) || string(char) == decimalSeparator:
			var number string
			number, i = makeNumber(textInput, i)

			if number != "" {
				tokenList = append(tokenList, Token{NUMBER, number})
			}

			i -= 1

		case unicode.IsLetter(rune(char)):
			var identifier string
			identifier, i = makeIdentifier(textInput, i)

			var token = Token{IDENTIFIER, identifier}
			for _, v := range defined_identifiers.Keywords {
				if v == identifier {
					token = Token{KEYWORD, identifier}
				}
			}

			tokenList = append(tokenList, token)
			i -= 1

		default:
			return tokenList, errors.InvalidCharacterError{Character: char, Position: i}
		}

		i++
	}

	return tokenList, nil
}

func makeNumber(textInput string, index int) (string, int) {
	var number = string(textInput[index])

	var decimalPart = false
	var i = index + 1

	if number == decimalSeparator {
		decimalPart = true
	}

	for i < len(textInput) {
		var char = string(textInput[i])

		if decimalPart && char == decimalSeparator {
			break
		}

		if char == decimalSeparator {
			decimalPart = true
		} else if !unicode.IsDigit(rune(char[0])) {
			break
		}

		number = number + char
		i++
	}

	number = strings.TrimSuffix(number, decimalSeparator)

	if len(number) > 0 && string(number[0]) == decimalSeparator {
		number = "0" + number
	}
	return number, i
}

func makeIdentifier(textInput string, index int) (string, int) {
	var identifier = string(textInput[index])

	var i = index + 1

	for i < len(textInput) {
		var char = textInput[i]

		if !unicode.IsLetter(rune(char)) && !unicode.IsDigit(rune(char)) {
			break
		}

		identifier = identifier + string(char)
		i++
	}

	return identifier, i
}
