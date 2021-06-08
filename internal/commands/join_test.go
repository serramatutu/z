package commands_test

import (
	"bytes"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestJoinNilSeparator(t *testing.T) {
	cmd := commands.NewJoin(nil, nil)
	_, err := cmd.Execute([][]byte{
		[]byte("aaa"),
		[]byte("bbb"),
	})

	if err == nil {
		t.Errorf("Join.Execute with nil separator should return error")
	}
}

func TestJoinSeparator(t *testing.T) {
	cmd := commands.NewJoin(nil, []byte(":"))
	result, err := cmd.Execute([][]byte{
		[]byte("aaa"),
		[]byte("bbb"),
	})
	if err != nil {
		t.Errorf("Unexpected error for Split.Execute")
	}

	expected := []byte("aaa:bbb")

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
