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

type ExtraJoinErr struct {
}

func (ExtraJoinErr) Error() string {
	return "join operation without matching split"
}

type ExtraSplitErr struct {
}

func (ExtraSplitErr) Error() string {
	return "split operation without matching join"
}
