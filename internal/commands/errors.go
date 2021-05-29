package commands

import "fmt"

type InvalidPositionalArgumentError struct {
	ArgumentName string
}

func (err InvalidPositionalArgumentError) Error() string {
	return fmt.Sprintf("invalid argument '%s'", err.ArgumentName)
}
