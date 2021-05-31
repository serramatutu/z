package argparse

import (
	"fmt"
	"strconv"
	"strings"
)

type InvalidRangeErr struct {
	Input string
}

func (err InvalidRangeErr) Error() string {
	return fmt.Sprintf("invalid range '%s'", err.Input)
}

func ParseRange(in string) (start int, end int, err error) {
	invalidErr := InvalidRangeErr{
		Input: in,
	}

	splitRange := strings.Split(in, ":")
	if len(splitRange) != 2 {
		err = invalidErr
		return
	}

	var range64 int64
	if splitRange[0] != "" {
		range64, err = strconv.ParseInt(splitRange[0], 10, 0)
		if err != nil {
			err = invalidErr
			return
		}
		start = int(range64)
	}

	if splitRange[1] != "" {
		range64, err = strconv.ParseInt(splitRange[1], 10, 0)
		if err != nil {
			err = invalidErr
			return
		}
		end = int(range64)
	}

	return
}
