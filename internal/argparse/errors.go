package argparse

import (
	"fmt"
	"strings"
)

type MissingPositionalArgumentsErr struct {
	ArgumentNames []string
}

func (err MissingPositionalArgumentsErr) Error() string {
	argumentNames := make([]string, len(err.ArgumentNames))
	for i, argName := range err.ArgumentNames {
		argumentNames[i] = fmt.Sprintf("\"%s\"", argName)
	}
	joinedNames := strings.Join(argumentNames, ", ")
	return fmt.Sprintf("missing positional arguments: %s", joinedNames)
}

type ExtraPositionalArgumentErr struct {
	ArgumentValue string
}

func (err ExtraPositionalArgumentErr) Error() string {
	return fmt.Sprintf("unused argument \"%s\"", err.ArgumentValue)
}

type InvalidPositionalArgumentErr struct {
	ArgumentName  string
	ArgumentValue string
}

func (err InvalidPositionalArgumentErr) Error() string {
	return fmt.Sprintf("invalid value \"%s\" for %s", err.ArgumentValue, err.ArgumentName)
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
