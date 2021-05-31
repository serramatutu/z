package commands

import (
	"fmt"
	"strings"
)

type MissingPositionalArgumentsErr struct {
	ArgumentNames []string
}

func (err MissingPositionalArgumentsErr) Error() string {
	return fmt.Sprintf("missing positional arguments %s", strings.Join(err.ArgumentNames, ", "))
}

type ExtraPositionalArgumentErr struct {
	ArgumentValue string
}

func (err ExtraPositionalArgumentErr) Error() string {
	return fmt.Sprintf("extra argument '%s'", err.ArgumentValue)
}

type InvalidPositionalArgumentErr struct {
	ArgumentName  string
	ArgumentValue string
}

func (err InvalidPositionalArgumentErr) Error() string {
	return fmt.Sprintf("invalid value '%s' for %s", err.ArgumentName, err.ArgumentValue)
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
