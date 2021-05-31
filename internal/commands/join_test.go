package commands_test

import (
	"bytes"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestJoinNilSeparator(t *testing.T) {
	cmd := commands.Join{}
	result, err := cmd.Execute([][]byte{
		[]byte("aaa"),
		[]byte("bbb"),
	})

	if err != nil {
		t.Errorf("Unexpected error for Join.Execute with nil separator")
	}

	expected := []byte("aaabbb")
	if !bytes.Equal(result, expected) {
		t.Errorf("Join with nil separator should concatenate (implicit join)")
	}
}

func TestJoinSeparator(t *testing.T) {
	cmd := commands.Join{
		Separator: []byte(":"),
	}
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
