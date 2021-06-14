package commands

import (
	"reflect"
	"regexp"
	"sort"
	"testing"
)

var sortInput = [][]byte{
	[]byte("a,1,6"),
	[]byte("a,2,6"),
	[]byte("b,1,6"),
	[]byte("c,2,6"),
	[]byte("c,3,6"),
	[]byte("d,3,6"),
}

func TestSortNilSeparator(t *testing.T) {
	sorter := newBytesKeySorter(sortInput, nil, 0)
	sort.Sort(sorter)

	if !reflect.DeepEqual(sorter.Data, sortInput) {
		t.Errorf("Invalid sorted result for bytesKeySorter with nil separator")
	}
}

func TestSortSeparator(t *testing.T) {
	sorter := newBytesKeySorter(sortInput, regexp.MustCompile(","), 0)
	// use stable to be deterministic
	sort.Stable(sorter)

	if !reflect.DeepEqual(sorter.Data, sortInput) {
		t.Errorf("Invalid sorted result for bytesKeySorter with nil separator")
	}
}
