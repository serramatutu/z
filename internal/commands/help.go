package commands

import (
	"errors"

	"github.com/serramatutu/z/help"
)

type Help struct {
	CommandName string
}

func (h Help) Err() error {
	return errors.New(help.Help[h.HelpFile()])
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

func (Help) Execute(in []byte) ([]byte, error) {
	return nil, nil
}

func ParseHelp(args []string) Help {
	var commandName string
	if len(args) > 0 {
		commandName = args[0]
	}

	return Help{
		CommandName: commandName,
	}
}
