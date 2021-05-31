package internal

import (
	"bytes"
	"container/list"
	"regexp"
	"strings"
	"testing"

	"github.com/serramatutu/z/internal/commands"
	"github.com/serramatutu/z/internal/config"
)

func TestExecuteSplitWithoutCommands(t *testing.T) {
	commandsList := list.New()
	commandsList.PushBack(commands.Split{})
	stop := commandsList.PushBack(commands.Join{})
	commandsList.PushBack(commands.Join{})

	result, lastRan, err := executeSplit([]byte("a\nb\nc\nd\ne"), commandsList.Front())
	if err != nil {
		t.Errorf("Unexpected error for executeSplit")
	}

	expected := []byte("abcde")
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected '%s' as executeSplit output but got '%s'", expected, result)
	}

	if lastRan != stop {
		t.Errorf("Expected executeSplit to consume exactly one join")
	}
}

func TestExecuteSplitWithCommands(t *testing.T) {
	commandsList := list.New()
	commandsList.PushBack(commands.Split{})
	commandsList.PushBack(commands.Length{})
	stop := commandsList.PushBack(commands.Join{})

	result, lastRan, err := executeSplit([]byte("a\nb\nc\nd\ne"), commandsList.Front())
	if err != nil {
		t.Errorf("Unexpected error for executeSplit")
	}

	expected := []byte("11111")
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected '%s' as executeSplit output but got '%s'", expected, result)
	}

	if lastRan != stop {
		t.Errorf("Expected executeSplit to consume exactly one join")
	}
}

func TestExecuteSplitNested(t *testing.T) {
	commandsList := list.New()
	sep, _ := regexp.Compile(":")
	commandsList.PushBack(commands.Split{
		Separator: sep,
	})
	sep, _ = regexp.Compile("b")
	commandsList.PushBack(commands.Split{
		Separator: sep,
	})
	commandsList.PushBack(commands.Length{})
	commandsList.PushBack(commands.Join{})
	stop := commandsList.PushBack(commands.Join{})

	result, lastRan, err := executeSplit([]byte("aba:aba:aba"), commandsList.Front())
	if err != nil {
		t.Errorf("Unexpected error for executeSplit")
	}

	expected := []byte("111111") // 6 'a' of length 1
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected '%s' as executeSplit output but got '%s'", expected, result)
	}

	if lastRan != stop {
		t.Errorf("Expected executeSplit to consume exactly two joins")
	}
}

func TestExecuteSplitImplicitJoin(t *testing.T) {
	commandsList := list.New()
	sep, _ := regexp.Compile(":")
	commandsList.PushBack(commands.Split{
		Separator: sep,
	})
	stop := commandsList.PushBack(commands.Length{})

	result, lastRan, err := executeSplit([]byte("a:a:a"), commandsList.Front())
	if err != nil {
		t.Errorf("Unexpected error for executeSplit")
	}

	expected := []byte("111")
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected '%s' as executeSplit output but got '%s'", expected, result)
	}

	if lastRan != stop {
		t.Errorf("Expected executeSplit to stop at last known command")
	}
}

func TestExecuteMapOnlyMapCommands(t *testing.T) {
	commandsList := list.New()
	commandsList.PushBack(commands.Length{})
	stop := commandsList.PushBack(commands.Length{})

	result, lastRan, err := executeMap([]byte("abcde"), commandsList.Front())
	if err != nil {
		t.Errorf("Unexpected error for executeMap")
	}

	expected := []byte("1")
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected '%s' as executeMap output but got '%s'", expected, result)
	}

	if lastRan != stop {
		t.Errorf("Expected executeMap to stop at last available MapCommand")
	}
}

func TestExecuteMapWithSplitCommand(t *testing.T) {
	commandsList := list.New()
	stop := commandsList.PushBack(commands.Length{})
	commandsList.PushBack(commands.Split{})
	commandsList.PushBack(commands.Length{})

	result, lastRan, err := executeMap([]byte("abcde"), commandsList.Front())
	if err != nil {
		t.Errorf("Unexpected error for executeMap")
	}

	expected := []byte("5")
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected '%s' as executeMap output but got '%s'", expected, result)
	}

	if lastRan != stop {
		t.Errorf("Expected executeMap to stop at last available MapCommand")
	}
}

func TestConfigExecuteExtraJoin(t *testing.T) {
	c := config.Config{
		Err:      nil,
		Commands: list.New(),
	}
	c.Commands.PushBack(commands.Length{})
	c.Commands.PushBack(commands.Join{})

	_, err := executeConfig(c, []byte("abcde"))
	if err == nil {
		t.Errorf("Expected 'ExtraJoinErr' when join is missing but got nil")
	}
}

func TestConfigExecuteOk(t *testing.T) {
	c := config.Config{
		Err:      nil,
		Commands: list.New(),
	}
	sep, _ := regexp.Compile(":")
	c.Commands.PushBack(commands.Split{
		Separator: sep,
	})
	c.Commands.PushBack(commands.Length{})
	c.Commands.PushBack(commands.Join{})

	result, err := executeConfig(c, []byte("a:aa:a"))
	if err != nil {
		t.Errorf("Unexpected error for config.Config.Execute")
	}

	expected := []byte("121")
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected '%s' as config.Config.Execute output but got '%s'", expected, result)
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
