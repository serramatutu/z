package commands

import (
	"errors"
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
	if r.Target == nil {
		return nil, errors.New("Replace target cannot be nil")
	}

	if r.RangeStart == 0 && r.RangeEnd == 0 {
		return r.Target.ReplaceAll(in, r.Replacement), nil
	}

	if r.RangeStart == r.RangeEnd {
		out := make([]byte, len(in))
		copy(out, in)
		return out, nil
	}

	// Replace manually

	allLocations := r.Target.FindAllIndex(in, -1)
	start := len(in)
	last := 0

	if len(allLocations) > r.RangeStart {
		start = r.RangeStart
	}

	if r.RangeEnd > 0 {
		if r.RangeEnd >= len(allLocations) {
			last = len(allLocations) - 1
		} else {
			last = r.RangeEnd - 1
		}
	} else {
		last = len(allLocations) - 1 + r.RangeEnd
	}

	if start > last {
		out := make([]byte, len(in))
		copy(out, in)
		return out, nil
	}

	// calculate the total resulting size and allocate it
	// this ensures all append calls do not reallocate the
	// underlying array, which would waste lots of time
	totalRemovedSize := 0
	for _, location := range allLocations[start : last+1] {
		totalRemovedSize += location[1] - location[0]
	}
	totalSize := len(in) - totalRemovedSize + (last-start+1)*len(r.Replacement)
	out := make([]byte, totalSize)

	// copy everything until the replacement start
	out = out[:allLocations[start][0]]
	copy(out, in[:allLocations[start][0]])
	replacementBytes := []byte(r.Replacement)

	// for each replacement, append its bytes and copy everything
	// until the next
	for i := start; i < last; i++ {
		out = append(out, replacementBytes...)

		currLocation := allLocations[i][1]
		nextLocation := allLocations[i+1][0]
		out = append(out, in[currLocation:nextLocation]...)
	}
	// append the last replacement and the final bytes
	out = append(out, replacementBytes...)
	out = append(out, in[allLocations[last][1]:]...)

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
