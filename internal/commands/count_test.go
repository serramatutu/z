package commands_test

import (
	"bytes"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestCountExecute(t *testing.T) {
	c := commands.Count{}
	args := [][]byte{
		[]byte("abcd"),
		[]byte("efgh"),
		[]byte("ijkl"),
	}
	result, err := c.Execute(args)
	if err != nil {
		t.Errorf("Count.Execute should never return error")
	}

	if !bytes.Equal(result, []byte("3")) {
		t.Errorf("Count.Execute result is wrong")
	}
}
