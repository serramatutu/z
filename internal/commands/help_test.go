package commands

import "testing"

func TestHelpFileForNoCommand(t *testing.T) {
	help := Help{
		CommandName: "",
	}

	if help.HelpFile() != "z" {
		t.Errorf("Help command should display 'z' helpfile when there's no command")
	}
}

func TestHelpFileForInvalidCommand(t *testing.T) {
	help := Help{
		CommandName: "invalid",
	}

	if help.HelpFile() != "z" {
		t.Errorf("Help command should display 'z' helpfile when command name is invalid")
	}
}

func TestHelpFileForValidCommand(t *testing.T) {
	help := Help{
		CommandName: "length",
	}

	if help.HelpFile() != "length" {
		t.Errorf("Help command should display subcommand helpfile when it is valid")
	}
}

func TestParseHelpNoSubcommand(t *testing.T) {
	args := []string{}
	help := ParseHelp(args)

	if help.CommandName != "" {
		t.Errorf("ParseHelp should return empty command name when no args are given. Got '%s'", help.CommandName)
	}
}

func TestParseHelpSubcommand(t *testing.T) {
	args := []string{"subcommand"}
	help := ParseHelp(args)

	if help.CommandName != "subcommand" {
		t.Errorf("ParseHelp should return parse subcommand. Got '%s'", help.CommandName)
	}
}
