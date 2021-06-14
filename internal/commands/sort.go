package commands

import (
	"regexp"
	"sort"

	"github.com/serramatutu/z/internal/utility"
)

type Sort struct {
	err error

	FieldSeparator *regexp.Regexp
	Index          int
}

func (s Sort) Err() error {
	return s.err
}

func (Sort) Name() string {
	return "sort"
}

func (Sort) HelpFile() string {
	return "sort"
}

type bytesKeySorter struct {
	keys []string
	Data [][]byte
}

func newBytesKeySorter(data [][]byte, fieldSeparator *regexp.Regexp, index int) bytesKeySorter {
	b := bytesKeySorter{
		keys: make([]string, len(data)),
		Data: data,
	}

	for i, val := range data {
		b.keys[i] = utility.ExtractKeyFromFields(val, fieldSeparator, index)
	}

	return b
}

func (s bytesKeySorter) Len() int {
	return len(s.Data)
}

func (s bytesKeySorter) Less(i, j int) bool {
	return s.keys[i] < s.keys[j]
}

func (s bytesKeySorter) Swap(i, j int) {
	s.keys[i], s.keys[j] = s.keys[j], s.keys[i]
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}

func (s Sort) Execute(in [][]byte) ([][]byte, error) {
	out := make([][]byte, len(in))
	copy(out, in)

	sorter := newBytesKeySorter(out, s.FieldSeparator, s.Index)
	sort.Sort(sorter)

	return out, nil
}

func NewSort(err error, fieldSeparator *regexp.Regexp, index int) Sort {
	return Sort{
		err: err,

		FieldSeparator: fieldSeparator,
		Index:          index,
	}
}
