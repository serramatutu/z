package commands_test

import (
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestHelpFileForNoCommand(t *testing.T) {
	help := commands.Help{
		CommandName: "",
	}

	if help.HelpFile() != "z" {
		t.Errorf("Help command should display 'z' helpfile when there's no command")
	}
}

func TestHelpFileForInvalidCommand(t *testing.T) {
	help := commands.Help{
		CommandName: "invalid",
	}

	if help.HelpFile() != "z" {
		t.Errorf("Help command should display 'z' helpfile when command name is invalid")
	}
}

func TestHelpFileForValidCommand(t *testing.T) {
	help := commands.Help{
		CommandName: "length",
	}

	if help.HelpFile() != "length" {
		t.Errorf("Help command should display subcommand helpfile when it is valid")
	}
}
