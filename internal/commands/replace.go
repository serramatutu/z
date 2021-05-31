package commands

import (
	"fmt"
	"regexp"
	// "github.com/serramatutu/z/internal/argparse"
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
	if r.RangeStart == 0 && r.RangeEnd == 0 {
		return r.Target.ReplaceAll(in, r.Replacement), nil
	}

	out := make([]byte, len(in))
	if r.RangeStart == r.RangeEnd {
		copy(out, in)
		return out, nil
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
		copy(out, in)
		return out, nil
	}

	// TODO: optimize
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

func NewReplace(err error, target *regexp.Regexp, replacement []byte, rangeStart, rangeEnd int) Replace {
	return Replace{
		err:         err,
		Target:      target,
		Replacement: replacement,
		RangeStart:  rangeStart,
		RangeEnd:    rangeEnd,
	}
}
