package commands

import (
	"fmt"
)

type Length struct {
	arg string
}

func (l Length) Err() error {
	if l.arg != "" {
		return InvalidPositionalArgumentErr{
			ArgumentName: l.arg,
		}
	}

	return nil
}

func (Length) Name() string {
	return "length"
}

func (Length) HelpFile() string {
	return "length"
}

func (Length) Execute(in string) (string, error) {
	return fmt.Sprint(len(in)), nil
}

func ParseLength(args []string) Length {
	if len(args) > 0 {
		return Length{
			arg: args[0],
		}
	}

	return Length{}
}
