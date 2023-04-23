package errors

import "fmt"

type InvalidCharacter struct {
	Character uint8
	Position  int
}

func (err InvalidCharacter) Error() string {
	return fmt.Sprintf("invalid character '%c' at position %d", err.Character, err.Position)
}
