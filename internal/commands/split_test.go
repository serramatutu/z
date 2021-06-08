package commands_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestSplitNilSeparator(t *testing.T) {
	cmd := commands.NewSplit(nil, nil)
	_, err := cmd.Execute([]byte("aaa:bbb-ccc\nddd_eee"))
	if err == nil {
		t.Errorf("Split.Execute with nil separator should return error")
	}
}

func TestSplitSeparator(t *testing.T) {
	cmd := commands.NewSplit(nil, regexp.MustCompile(":"))
	result, err := cmd.Execute([]byte("aaa:bbb-ccc\nddd_eee"))
	if err != nil {
		t.Errorf("Unexpected error for Split.Execute")
	}

	expected := [][]byte{
		[]byte("aaa"),
		[]byte("bbb-ccc\nddd_eee"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
