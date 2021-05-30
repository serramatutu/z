package commands

import (
	"fmt"
)

type InvalidPositionalArgumentErr struct {
	ArgumentName string
}

func (err InvalidPositionalArgumentErr) Error() string {
	return fmt.Sprintf("invalid argument '%s'", err.ArgumentName)
}

type InvalidPipeErr struct {
}

func (InvalidPipeErr) Error() string {
	return "invalid pipe"
}
