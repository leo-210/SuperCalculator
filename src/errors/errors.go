package errors

import "fmt"

type InvalidCharacterError struct {
	Character uint8
	Position  int
}

func (err InvalidCharacterError) Error() string {
	return fmt.Sprintf("invalid character '%c' at position %d", err.Character, err.Position)
}

type ExpressionExpectedError struct {
	Position int
}

func (err ExpressionExpectedError) Error() string {
	return fmt.Sprintf("expression expected at position %d", err.Position)
}

type MissingClosingParenthesisError struct {
	Position int
}

func (err MissingClosingParenthesisError) Error() string {
	return fmt.Sprintf("missing closing parenthesis at position %d", err.Position)
}

type SyntaxError struct {
	Message  string
	Position int
}

func (err SyntaxError) Error() string {
	return fmt.Sprintf("%s at position %d", err.Message, err.Position)
}

type MathError struct {
	Message string
}

func (err MathError) Error() string {
	return err.Message
}

type NotImplementedYetError struct{}

func (err NotImplementedYetError) Error() string {
	return "feature not fully implemented yet"
}
