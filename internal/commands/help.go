package commands

import "errors"

type Help struct {
}

func (l Help) Err() error {
	return errors.New("")
}

func (Help) Name() string {
	return "help"
}

func (Help) Execute(in string) (string, error) {
	return "", nil
}

func ParseHelp(args []string) Help {
	return Help{}
}
