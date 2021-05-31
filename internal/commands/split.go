package commands

import (
	"bytes"
)

type Split struct {
	err error

	Separator []byte
}

func (s Split) Err() error {
	return s.err
}

func (Split) Name() string {
	return "split"
}

func (Split) HelpFile() string {
	return "split"
}

func (s Split) Execute(in []byte) ([][]byte, error) {
	return bytes.Split(in, s.Separator), nil
}

func ParseSplit(args []string) Split {
	var err error
	var sep []byte

	switch len(args) {
	case 0:
		sep = []byte("\n")
	case 1:
		sep = []byte(args[0])
	default:
		err = ExtraPositionalArgumentErr{
			ArgumentValue: args[1],
		}
		sep = nil
	}

	return Split{
		err:       err,
		Separator: sep,
	}
}
