package commands

import (
	"bytes"
	"errors"
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
	if j.err != nil {
		return nil, j.err
	}

	return bytes.Join(in, j.Separator), nil
}

func NewJoin(err error, separator []byte) Join {
	if err == nil && separator == nil {
		err = errors.New("cannot execute join with nil separator")
	}

	return Join{
		err:       err,
		Separator: separator,
	}
}
