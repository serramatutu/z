package commands

import (
	"github.com/serramatutu/z/help"
)

type Help struct {
	CommandName string
}

func (h Help) Err() error {
	return nil
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

func (h Help) Execute() []byte {
	return []byte(help.Help[h.HelpFile()])
}
