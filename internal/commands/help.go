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
	return h.CommandName
}

func (Help) Execute(in string) (string, error) {
	return "", nil
}

func ParseHelp(args []string) Help {
	commandName := "z"
	if len(args) > 0 {
		if _, ok := help.Help[args[0]]; ok {
			commandName = args[0]
		}
	}

	return Help{
		CommandName: commandName,
	}
}
