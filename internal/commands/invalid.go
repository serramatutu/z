package commands

import (
	"fmt"
)

type Invalid struct {
	CommandName string
}

func (i Invalid) Err() error {
	return fmt.Errorf("Invalid command '%s'", i.CommandName)
}

func (i Invalid) Name() string {
	return i.CommandName
}

func (Invalid) Execute(str string) (string, error) {
	return "", nil
}
