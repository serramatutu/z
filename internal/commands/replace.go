package commands

import (
	"fmt"
)

type Replace struct {
	err error

	Target      string
	Replacement string
	RangeStart  int
	RangeEnd    int
}

func (r Replace) Err() error {
	return r.err
}

func (Replace) Name() string {
	return "replace"
}

func (Replace) HelpFile() string {
	return "replace"
}

func (Replace) Execute(in []byte) ([]byte, error) {
	return []byte(fmt.Sprint(len(in))), nil
}

func ParseReplace(args []string) Replace {
	var err error

	switch len(args) {
	case 0:
		err = MissingPositionalArgumentsErr{
			ArgumentNames: []string{"string-or-pattern", "replace-string"},
		}
	case 1:
		err = MissingPositionalArgumentsErr{
			ArgumentNames: []string{"replace-string"},
		}
	case 2:
	case 3:
	default:
		err = ExtraPositionalArgumentErr{
			ArgumentName: args[3],
		}
	}

	return Replace{
		err: err,
	}
}
