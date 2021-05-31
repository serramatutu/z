package commands

import (
	"bytes"
)

type Split struct {
	extraArg string

	Separator []byte
}

func (s Split) Err() error {
	if s.extraArg != "" {
		return InvalidPositionalArgumentErr{
			ArgumentName: s.extraArg,
		}
	}

	return nil
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
	extraArg := ""
	var sep []byte

	switch len(args) {
	case 0:
		sep = []byte("\n")
	case 1:
		sep = []byte(args[0])
	default:
		extraArg = args[1]
		sep = nil
	}

	return Split{
		extraArg:  extraArg,
		Separator: sep,
	}
}
