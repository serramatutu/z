package commands_test

import (
	"bytes"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestLengthExecuteNoMode(t *testing.T) {
	l := commands.NewLength(nil, "")
	_, err := l.Execute([]byte("1234"))
	if err == nil {
		t.Errorf("Length.Execute should return error when no mode supplied")
	}
}

func TestLengthExecuteModeBytes(t *testing.T) {
	l := commands.NewLength(nil, commands.Bytes)
	result, err := l.Execute([]byte("ç"))
	if err != nil {
		t.Errorf("Unexpected error for Length.Execute with valid mode")
	}

	if !bytes.Equal(result, []byte("2")) {
		t.Errorf("Length.Execute result is wrong")
	}
}

func TestLengthExecuteModeUnicode(t *testing.T) {
	l := commands.NewLength(nil, commands.Unicode)
	result, err := l.Execute([]byte("ç"))
	if err != nil {
		t.Errorf("Unexpected error for Length.Execute with valid mode")
	}

	if !bytes.Equal(result, []byte("1")) {
		t.Errorf("Length.Execute result is wrong")
	}
}
