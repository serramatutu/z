package commands

import (
	"regexp"
)

type Split struct {
	err error

	Separator *regexp.Regexp
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
	sep := s.Separator
	if sep == nil {
		sep = regexp.MustCompile("\n")
	}

	split := sep.Split(string(in), -1)
	out := make([][]byte, len(split))
	for i, val := range split {
		out[i] = []byte(val)
	}

	return out, nil
}

func NewSplit(err error, separator *regexp.Regexp) Split {
	return Split{
		err:       err,
		Separator: separator,
	}
}
