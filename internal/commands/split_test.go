package commands_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestNilSeparator(t *testing.T) {
	cmd := commands.Split{}
	result, err := cmd.Execute([]byte("aaa:bbb-ccc\nddd_eee"))
	if err != nil {
		t.Errorf("Unexpected error for Split.Execute with nil separator")
	}

	expected := [][]byte{
		[]byte("aaa:bbb-ccc"),
		[]byte("ddd_eee"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Split should use '\\n' as default separator")
	}
}

func TestSplit(t *testing.T) {
	cmd := commands.Split{
		Separator: regexp.MustCompile(":"),
	}
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
