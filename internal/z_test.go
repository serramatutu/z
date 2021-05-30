package internal

import (
	"bytes"
	"container/list"
	"strings"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestExecuteMapOnlyMapCommands(t *testing.T) {
	commandsList := list.New()
	commandsList.PushBack(commands.Length{})
	commandsList.PushBack(commands.Length{})

	result, lastRan, err := executeMap([]byte("abcde"), commandsList.Front())
	if err != nil {
		t.Errorf("Unexpected error for executeMap")
	}

	expected := []byte("5")
	if !bytes.Equal(result, []byte("1")) {
		t.Errorf("Expected '%s' as executeMap output but got '%s'", expected, result)
	}

	if lastRan != commandsList.Back() {
		t.Errorf("Expected executeMap to run until end of list")
	}
}

func TestExecuteMapWithSplitCommand(t *testing.T) {
	commandsList := list.New()
	commandsList.PushBack(commands.Length{})
	commandsList.PushBack(commands.Split{})
	commandsList.PushBack(commands.Length{})

	result, lastRan, err := executeMap([]byte("abcde"), commandsList.Front())
	if err != nil {
		t.Errorf("Unexpected error for executeMap")
	}

	expected := []byte("5")
	if !bytes.Equal(result, []byte("5")) {
		t.Errorf("Expected '%s' as executeMap output but got '%s'", expected, result)
	}

	if lastRan != commandsList.Front() {
		t.Errorf("Expected executeMap to run until last map command")
	}
}

func TestWriteLength(t *testing.T) {
	args := []string{"z", "length"}
	in := strings.NewReader("1234\n\n")

	var out bytes.Buffer
	Z(args, in, &out)

	expected := []byte("6")

	if !bytes.Equal(out.Bytes(), []byte("6")) {
		t.Errorf("Expected '%s' as Z output but got '%s'", expected, out.String())
	}
}
