package commands_test

import (
	"bytes"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestLengthExecute(t *testing.T) {
	l := commands.Length{}
	result, err := l.Execute([]byte("1234"))
	if err != nil {
		t.Errorf("Length.Execute should never return error")
	}

	if !bytes.Equal(result, []byte("4")) {
		t.Errorf("Length.Execute result is wrong")
	}
}
