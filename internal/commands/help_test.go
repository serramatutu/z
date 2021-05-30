package commands

import "testing"

func TestHelpFileForNoCommand(t *testing.T) {
	cmd := Help{
		CommandName: "",
	}

	if cmd.HelpFile() != "z" {
		t.Errorf("Help command should display 'z' helpfile when there's no command")
	}
}

func TestHelpFileForInvalidCommand(t *testing.T) {
	cmd := Help{
		CommandName: "invalid",
	}

	if cmd.HelpFile() != "z" {
		t.Errorf("Help command should display 'z' helpfile when command name is invalid")
	}
}

func TestHelpFileForValidCommand(t *testing.T) {
	cmd := Help{
		CommandName: "length",
	}

	if cmd.HelpFile() != "length" {
		t.Errorf("Help command should display subcommand helpfile when it is valid")
	}
}
