package commands

import (
	"fmt"
)

type Invalid struct {
	CommandName string
}

func (i Invalid) Err() error {
	return fmt.Errorf("invalid command '%s'", i.CommandName)
}

func (i Invalid) Name() string {
	return i.CommandName
}

func (Invalid) HelpFile() string {
	return "z"
}

func (Invalid) Execute(in []byte) ([]byte, error) {
	return nil, nil
}
