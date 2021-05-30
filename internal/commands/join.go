package commands

import (
	"bytes"
)

type Join struct {
	extraArg string

	Separator []byte
}

func (j Join) Err() error {
	if j.extraArg != "" {
		return InvalidPositionalArgumentErr{
			ArgumentName: j.extraArg,
		}
	}

	return nil
}

func (Join) Name() string {
	return "join"
}

func (Join) HelpFile() string {
	return "join"
}

func (j Join) Execute(in [][]byte) ([]byte, error) {
	return bytes.Join(in, j.Separator), nil
}

func ParseJoin(args []string) Join {
	extraArg := ""
	var sep []byte

	switch len(args) {
	case 0:
		sep = []byte("")
	case 1:
		sep = []byte(args[0])
	default:
		extraArg = args[1]
		sep = nil
	}

	return Join{
		extraArg:  extraArg,
		Separator: sep,
	}
}
