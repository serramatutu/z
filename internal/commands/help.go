package commands

import (
	"errors"

	"github.com/serramatutu/z/help"
)

type Help struct {
	CommandName string
}

func (l Help) Err() error {
	return errors.New("")
}

func (h Help) Name() string {
	return "help"
}

func (h Help) HelpFile() string {
	if _, ok := help.Help[h.CommandName]; ok {
		return h.CommandName
	}
	return "z"
}

func (Help) Execute(in string) (string, error) {
	return "", nil
}

func ParseHelp(args []string) Help {
	commandName := "z"
	if len(args) > 0 {
		commandName = args[0]
	}

	return Help{
		CommandName: commandName,
	}
}
