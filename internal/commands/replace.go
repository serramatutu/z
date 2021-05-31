package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Replace struct {
	err error

	Target      *regexp.Regexp
	Replacement []byte
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

// FIXME
func (r Replace) Execute(in []byte) ([]byte, error) {
	if r.RangeStart == r.RangeEnd {
		if r.RangeStart == 0 {
			return r.Target.ReplaceAll(in, r.Replacement), nil
		}
		return in, nil
	}

	// Replace manually

	allLocations := r.Target.FindAllIndex(in, -1)
	start := len(in)
	end := 0

	if len(allLocations) > r.RangeStart {
		start = r.RangeStart
	}

	switch {
	case r.RangeEnd == 0:
		end = len(allLocations) - 1
	case r.RangeEnd < 0:
		end = len(allLocations) - 1 + r.RangeEnd
	}

	fmt.Printf("start: %v, end: %v, matches: %v\n", start, end, len(allLocations))

	if start > end {
		return in, nil
	}

	// TODO: optimize
	out := make([]byte, len(in))
	copy(out[:allLocations[start][0]], in[:allLocations[start][0]])
	replacementBytes := []byte(r.Replacement)
	for i := start; i < end; i++ {
		out = append(out, replacementBytes...)

		currLocation := allLocations[i][1]
		nextLocation := allLocations[i+1][0]
		out = append(out, in[currLocation:nextLocation]...)
	}
	out = append(out, replacementBytes...)
	out = append(out, in[allLocations[end][1]:]...)

	return out, nil
}

func ParseReplace(args []string) Replace {
	var err error
	var target *regexp.Regexp
	var replacement []byte
	var rangeStart, rangeEnd int

	switch len(args) {
	case 0:
		err = MissingPositionalArgumentsErr{
			ArgumentNames: []string{"pattern", "replace-string"},
		}
	case 1:
		err = MissingPositionalArgumentsErr{
			ArgumentNames: []string{"replace-string"},
		}
	case 3:
		invalidRangeErr := InvalidPositionalArgumentErr{
			ArgumentName:  "occurrence-range",
			ArgumentValue: args[2],
		}

		splitRange := strings.Split(args[2], ":")
		if len(splitRange) != 2 {
			err = invalidRangeErr
			break
		}

		var range64 int64
		if splitRange[0] != "" {
			range64, err = strconv.ParseInt(splitRange[0], 10, 0)
			if err != nil {
				err = invalidRangeErr
				break
			}
			rangeStart = int(range64)
		}

		if splitRange[1] != "" {
			range64, err = strconv.ParseInt(splitRange[1], 10, 0)
			if err != nil {
				err = invalidRangeErr
				break
			}
			rangeEnd = int(range64)
		}

		fallthrough
	case 2:
		target, err = regexp.Compile(args[0])
		if err != nil {
			err = InvalidPositionalArgumentErr{
				ArgumentName:  "pattern",
				ArgumentValue: args[0],
			}
		}
		replacement = []byte(args[1])
	default:
		err = ExtraPositionalArgumentErr{
			ArgumentValue: args[3],
		}
	}

	return Replace{
		err: err,

		Target:      target,
		Replacement: replacement,
		RangeStart:  rangeStart,
		RangeEnd:    rangeEnd,
	}
}
