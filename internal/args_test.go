package internal

import (
	"reflect"
	"strings"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

var commandTypeMap = map[string]interface{}{
	"invalid": commands.Invalid{},

	"help":   commands.Help{},
	"length": commands.Length{},
}

func TestParseCommandTypes(t *testing.T) {
	arr := make([]string, 1)
	for cmdName, expectedCmd := range commandTypeMap {
		t.Run(cmdName, func(t *testing.T) {
			arr[0] = cmdName
			cmd := parseCommand(arr)

			cmdType := reflect.TypeOf(cmd)
			expectedType := reflect.TypeOf(expectedCmd)

			if cmdType != expectedType {
				t.Errorf("Expected type '%s' for command '%s' but got '%s'", expectedType, cmdName, cmdType)
			}
		})
	}
}

func TestParseArgsSingleCommand(t *testing.T) {
	args := []string{"z", "length"}
	config := parseArgs(args)

	if config.Commands == nil {
		t.Errorf("Expected command list of size 1 but got nil")
		return
	}

	if config.Commands.Len() != 1 {
		t.Errorf("Expected command list of size 1 but got %v", config.Commands.Len())
		return
	}

	command := config.Commands.Front().Value.(commands.Command)
	switch command.(type) {
	case commands.Length:
	default:
		t.Errorf("Expected 'Length' command but got %s", reflect.TypeOf(command))
	}
}

func TestParseArgsChainedCommands(t *testing.T) {
	args := []string{"z", "length", "_", "length"}
	config := parseArgs(args)

	if config.Commands == nil {
		t.Errorf("Expected command list of size 2 but got nil")
		return
	}

	if config.Commands.Len() != 2 {
		t.Errorf("Expected command list of size 2 but got %v", config.Commands.Len())
		return
	}

	for e := config.Commands.Front(); e != nil; e = e.Next() {
		command := e.Value.(commands.Command)
		switch command.(type) {
		case commands.Length:
		default:
			t.Errorf("Expected 'Length' command but got %s", reflect.TypeOf(command))
		}
	}
}

func TestParseArgsInvalidPipeChain(t *testing.T) {
	invalidChains := [][]string{
		{"z", "_", "length"},
		{"z", "_", "length", "_"},
		{"z", "length", "_"},
		{"z", "length", "_", "_", "length"},
	}

	expectedType := reflect.TypeOf(commands.InvalidPipeErr{})

	for _, args := range invalidChains {
		t.Run(strings.Join(args, ","), func(t *testing.T) {
			config := parseArgs(args)

			if config.Commands != nil {
				t.Errorf("Expected nil command list of size but got pointer")
				return
			}

			configErrType := reflect.TypeOf(config.Err)

			if configErrType != expectedType {
				t.Errorf("Expected '%s' but got '%s'", expectedType, configErrType)
			}
		})
	}
}

func TestParseArgsNoCommand(t *testing.T) {
	args := []string{"z"}
	config := parseArgs(args)

	if config.Err == nil {
		t.Errorf("Expected 'ArgumentErr' but got nil")
		return
	}
}
