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
	sep := j.Separator
	if sep == nil {
		sep = []byte("")
	}

	return bytes.Join(in, sep), nil
}

func NewJoin(err error, separator []byte) Join {
	return Join{
		err:       err,
		Separator: separator,
	}
}
