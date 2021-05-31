package commands

import (
	"fmt"
)

type Length struct {
	err error
}

func (l Length) Err() error {
	return l.err
}

func (Length) Name() string {
	return "length"
}

func (Length) HelpFile() string {
	return "length"
}

func (Length) Execute(in []byte) ([]byte, error) {
	return []byte(fmt.Sprint(len(in))), nil
}

func ParseLength(args []string) Length {
	var err error
	if len(args) > 0 {
		err = ExtraPositionalArgumentErr{
			ArgumentName: args[0],
		}
	}

	return Length{
		err: err,
	}
}
