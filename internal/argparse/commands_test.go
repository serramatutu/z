package argparse_test

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/argparse"
)

func TestParseHelpNoSubcommand(t *testing.T) {
	args := []string{}
	help := argparse.ParseHelp(args)

	if help.CommandName != "" {
		t.Errorf("ParseHelp should return empty command name when no args are given. Got '%s'", help.CommandName)
	}
}

func TestParseHelpSubcommand(t *testing.T) {
	args := []string{"subcommand"}
	help := argparse.ParseHelp(args)

	if help.CommandName != "subcommand" {
		t.Errorf("ParseHelp should return parse subcommand. Got '%s'", help.CommandName)
	}
}

func TestParseJoinNoArgs(t *testing.T) {
	args := []string{}
	join := argparse.ParseJoin(args)

	if join.Err() != nil {
		t.Errorf("Unexpected error for ParseJoin with no args")
	}

	if join.Separator != nil {
		t.Errorf("ParseJoin should produce nil separator if not provided")
	}
}

func TestParseJoinSeparator(t *testing.T) {
	args := []string{":"}
	join := argparse.ParseJoin(args)

	if join.Err() != nil {
		t.Errorf("Unexpected error for ParseJoin with separator arg")
	}

	if join.Separator == nil {
		t.Errorf("ParseJoin should produce separator when it is provided")
	}

	expected := []byte(":")
	if !bytes.Equal(join.Separator, expected) {
		t.Errorf("Expected '%s' but got '%s' as separator for ParseJoin", expected, join.Separator)
	}
}

func TestParseJoinTooManyArgs(t *testing.T) {
	args := []string{":", "invalid"}
	join := argparse.ParseJoin(args)

	switch join.Err().(type) {
	case argparse.ExtraPositionalArgumentErr:
	default:
		t.Errorf("ParseJoin should return ExtraPositionalArgumentErr when there are invalid arguments")
	}
}

func TestParseLengthNoArgs(t *testing.T) {
	args := []string{}
	length := argparse.ParseLength(args)

	if length.Err() != nil {
		t.Errorf("ParseLength should not return error when no args are given")
	}
}

func TestParseLengthWithArgs(t *testing.T) {
	args := []string{"arg"}
	length := argparse.ParseLength(args)

	if length.Err() == nil {
		t.Errorf("ParseLength should return error when args are given")
	}
}

func TestParseReplaceMissingArgs(t *testing.T) {
	argsList := [][]string{
		{},
		{"notenough"},
	}

	for i, args := range argsList {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			replace := argparse.ParseReplace(args)
			switch replace.Err().(type) {
			case argparse.MissingPositionalArgumentsErr:
			default:
				t.Errorf("ParseReplace should fail when there are missing arguments.")
			}
		})
	}
}

func TestParseReplaceTooManyArgs(t *testing.T) {
	args := []string{":", "-", ":", "invalid"}
	replace := argparse.ParseReplace(args)

	switch replace.Err().(type) {
	case argparse.ExtraPositionalArgumentErr:
	default:
		t.Errorf("ParseJoin should return ExtraPositionalArgumentErr when there are invalid arguments")
	}
}

func TestParseReplaceNoRange(t *testing.T) {
	args := []string{":", "-"}
	replace := argparse.ParseReplace(args)

	if replace.Err() != nil {
		t.Errorf("Unexpected error for ParseReplace with pattern and replacement args")
	}

	expectedTarget := regexp.MustCompile(":")
	if !reflect.DeepEqual(*replace.Target, *expectedTarget) {
		t.Errorf("Invalid regex target for ParseReplace")
	}

	expectedReplacement := []byte("-")
	if !bytes.Equal(replace.Replacement, expectedReplacement) {
		t.Errorf("Expected '%s' but got '%s' as replacement for ParseReplacement", expectedReplacement, replace.Replacement)
	}

	if replace.RangeStart != 0 || replace.RangeEnd != 0 {
		t.Errorf("Expected 0:0 but got %v:%v as range for ParseReplacement", replace.RangeStart, replace.RangeEnd)
	}
}

func TestParseReplaceRange(t *testing.T) {
	args := []string{":", "-", "1:2"}
	replace := argparse.ParseReplace(args)

	if replace.RangeStart != 1 || replace.RangeEnd != 2 {
		t.Errorf("Expected 1:2 but got %v:%v as range for ParseReplacement", replace.RangeStart, replace.RangeEnd)
	}
}

// Range parsing tests are in "range_test.go"
