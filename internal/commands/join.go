package commands

import (
	"bytes"
)

type Join struct {
	err error

	Separator []byte
}

func (j Join) Err() error {
	return j.err
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
	var err error
	var sep []byte

	switch len(args) {
	case 0:
		sep = []byte("")
	case 1:
		sep = []byte(args[0])
	default:
		err = ExtraPositionalArgumentErr{
			ArgumentName: args[1],
		}
		sep = nil
	}

	return Join{
		err:       err,
		Separator: sep,
	}
}
