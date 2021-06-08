package commands

import (
	"errors"
	"fmt"
)

type LengthMode string

const (
	Unicode = LengthMode("unicode")
	Bytes   = LengthMode("bytes")
)

func (m LengthMode) Validate() error {
	switch m {
	case Unicode, Bytes:
		return nil
	}
	return errors.New("Invalid length mode")
}

type Length struct {
	err error

	Mode LengthMode
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

func (l Length) Execute(in []byte) ([]byte, error) {
	if l.err != nil {
		return nil, l.err
	}

	var length int
	switch l.Mode {
	case Bytes:
		length = len(in)
	case Unicode:
		length = len([]rune(string(in)))
	}
	return []byte(fmt.Sprint(length)), nil
}

func NewLength(err error, mode LengthMode) Length {
	if err == nil {
		err = mode.Validate()
	}

	return Length{
		err:  err,
		Mode: mode,
	}
}
