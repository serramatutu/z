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

func NewLength(err error) Length {
	return Length{
		err: err,
	}
}
