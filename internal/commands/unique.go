package commands

import (
	"regexp"

	"github.com/serramatutu/z/internal/utility"
)

type Unique struct {
	err error

	FieldSeparator *regexp.Regexp
	Index          int
}

func (u Unique) Err() error {
	return u.err
}

func (Unique) Name() string {
	return "unique"
}

func (Unique) HelpFile() string {
	return "unique"
}

func (u Unique) Execute(in [][]byte) ([][]byte, error) {
	found := make(map[string]bool)
	out := make([][]byte, len(in))[:0]

	for _, val := range in {
		key := utility.ExtractKeyFromFields(val, u.FieldSeparator, u.Index)

		if foundVal, ok := found[key]; !ok || !foundVal {
			out = append(out, val)
			found[key] = true
		}
	}

	return out, nil
}

func NewUnique(err error, fieldSeparator *regexp.Regexp, index int) Unique {
	return Unique{
		err: err,

		FieldSeparator: fieldSeparator,
		Index:          index,
	}
}
