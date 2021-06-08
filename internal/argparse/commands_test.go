package argparse_test

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/argparse"
	"github.com/serramatutu/z/internal/commands"
)

func TestParseCountNoArgs(t *testing.T) {
	args := []string{}
	count := argparse.ParseCount(args)

	if count.Err() != nil {
		t.Errorf("ParseCount should not return error when no args are given")
	}
}

func TestParseCountWithArgs(t *testing.T) {
	args := []string{"arg"}
	count := argparse.ParseCount(args)

	if count.Err() == nil {
		t.Errorf("ParseCount should return error when args are given")
	}
}

func TestParseHashNoAlgorithm(t *testing.T) {
	args := []string{}
	hash := argparse.ParseHash(args)

	if hash.Err() == nil {
		t.Errorf("ParseHash should return error when no algorithm is specified")
	}
}

func TestParseHashValidAlgorithm(t *testing.T) {
	validArgs := []string{"md5", "sha1", "sha224", "sha256"}
	for _, arg := range validArgs {
		t.Run(arg, func(t *testing.T) {
			hash := argparse.ParseHash([]string{arg})

			if hash.Err() != nil {
				t.Errorf("Unexpected error for ParseHash with valid arg")
			}

			if hash.Algorithm.Validate() != nil {
				t.Errorf("ParseHash with valid arg produced invalid HashAlgorithm enum value")
			}
		})
	}
}

func TestParseHashTooManyArgs(t *testing.T) {
	args := []string{"md5", "invalid"}
	hash := argparse.ParseHash(args)

	switch hash.Err().(type) {
	case argparse.ExtraPositionalArgumentErr:
	default:
		t.Errorf("ParseHash should return ExtraPositionalArgumentErr when there are invalid arguments")
	}
}

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

	switch length.Mode {
	case commands.Bytes:
	default:
		t.Errorf("Length default mode should be bytes")
	}
}

func TestParseLengthValidMode(t *testing.T) {
	validArgs := []string{"unicode", "bytes"}
	for _, arg := range validArgs {
		t.Run(arg, func(t *testing.T) {
			length := argparse.ParseLength([]string{arg})

			if length.Err() != nil {
				t.Errorf("Unexpected error for ParseLength with valid arg")
			}

			if length.Mode.Validate() != nil {
				t.Errorf("ParseLength with valid arg produced invalid LengthMode enum value")
			}
		})
	}
}

func TestParseMatchMissingArgs(t *testing.T) {
	args := []string{}

	match := argparse.ParseMatch(args)
	switch match.Err().(type) {
	case argparse.MissingPositionalArgumentsErr:
	default:
		t.Errorf("ParseMatch should fail when there are missing arguments.")
	}
}

func TestParseMatchOk(t *testing.T) {
	args := []string{":"}

	match := argparse.ParseMatch(args)
	if match.Err() != nil {
		t.Errorf("Unexpected error for ParseMatch with pattern arg")
	}

	if match.Pattern == nil {
		t.Errorf("ParseMatch should produce pattern when it is provided")
	}

	expected := regexp.MustCompile(":")
	if !reflect.DeepEqual(*match.Pattern, *expected) {
		t.Errorf("Invalid regex pattern for ParseMatch")
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

func TestParseSplitNoArgs(t *testing.T) {
	args := []string{}

	split := argparse.ParseSplit(args)
	if split.Err() != nil {
		t.Errorf("Unexpected error for ParseSplit with no args")
	}

	if split.Separator != nil {
		t.Errorf("ParseSplit should produce nil separator if not provided")
	}
}

func TestParseSplitSeparator(t *testing.T) {
	args := []string{":"}

	split := argparse.ParseSplit(args)
	if split.Err() != nil {
		t.Errorf("Unexpected error for ParseSplit with separator arg")
	}

	if split.Separator == nil {
		t.Errorf("ParseSplit should produce separator when it is provided")
	}

	expected := regexp.MustCompile(":")
	if !reflect.DeepEqual(*split.Separator, *expected) {
		t.Errorf("Invalid regex separator for ParseSplit")
	}
}
