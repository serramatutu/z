package argparse_test

import (
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
