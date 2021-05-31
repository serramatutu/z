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
	ArgumentName string
}

func (err ExtraPositionalArgumentErr) Error() string {
	return fmt.Sprintf("extra argument '%s'", err.ArgumentName)
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
