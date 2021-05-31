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
	split := s.Separator.Split(string(in), -1)
	out := make([][]byte, len(split))
	for i, val := range split {
		out[i] = []byte(val)
	}

	return out, nil
}

func ParseSplit(args []string) Split {
	var err error
	var sep *regexp.Regexp

	switch len(args) {
	case 0:
		sep, _ = regexp.Compile("\n")
	case 1:
		sep, err = regexp.Compile(args[0])
		if err != nil {
			err = InvalidPositionalArgumentErr{
				ArgumentName:  "pattern",
				ArgumentValue: args[0],
			}
		}
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
