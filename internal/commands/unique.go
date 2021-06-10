package commands

import (
	"regexp"
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

	wholeKey := false
	if u.FieldSeparator == nil {
		wholeKey = true
	}

	for _, val := range in {
		var key string
		if wholeKey {
			key = string(val)
		} else {
			separatorMatches := u.FieldSeparator.FindAllIndex(val, u.Index+1)
			if separatorMatches == nil {
				// consider the whole string as the first field
				if u.Index == 0 {
					key = string(val)
				}
				// field index is not available in this string, so key is empty
			} else if len(separatorMatches) >= u.Index {
				// field index is available, so find the field and use it as key
				keyStart := 0
				if u.Index > 0 {
					keyStart = separatorMatches[u.Index-1][1]
				}

				keyEnd := len(val)
				if u.Index < len(separatorMatches) {
					keyEnd = separatorMatches[u.Index][0]
				}

				keyBytes := val[keyStart:keyEnd]
				key = string(keyBytes)
			}

			// field index is not available in this string, so key is empty
		}

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
