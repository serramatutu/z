package internal

import (
	"container/list"
	"io"
	"io/ioutil"

	"github.com/serramatutu/z/internal/commands"
)

type Config struct {
	Err      error
	Commands *list.List
}

func (config Config) Execute(bytes []byte) ([]byte, error) {
	var err error
	for e := config.Commands.Front(); e != nil; e = e.Next() {
		command := e.Value.(commands.Command)
		bytes, err = command.Execute(bytes)
		if err != nil {
			return nil, err
		}
	}
	return bytes, nil
}

func Z(args []string, r io.Reader, w io.Writer) error {
	config := parseArgs(args)
	if config.Err != nil {
		return config.Err
	}

	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var output []byte
	output, err = config.Execute(contents)
	if err != nil {
		return err
	}

	w.Write(output)

	return nil
}
