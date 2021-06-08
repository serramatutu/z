package commands_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestNewMatchWithNilPattern(t *testing.T) {
	cmd := commands.NewMatch(nil, nil)
	if cmd.Err() == nil {
		t.Errorf("Expected NewMatch to result in error when pattern is nil")
	}
}

func TestMatch(t *testing.T) {
	cmd := commands.NewMatch(nil, regexp.MustCompile("[a-z]+"))
	result, err := cmd.Execute([]byte("aaa:bbb-ccc\nddd_eee"))
	if err != nil {
		t.Errorf("Unexpected error for Match.Execute")
	}

	expected := [][]byte{
		[]byte("aaa"),
		[]byte("bbb"),
		[]byte("ccc"),
		[]byte("ddd"),
		[]byte("eee"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
