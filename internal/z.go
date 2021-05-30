package internal

import (
	"bufio"
	"container/list"
	"fmt"
	"io"

	"github.com/serramatutu/z/internal/commands"
)

type Config struct {
	Err      error
	Commands *list.List
}

func (config Config) Execute(str string) (string, error) {
	var err error
	for e := config.Commands.Front(); e != nil; e = e.Next() {
		command := e.Value.(commands.Command)
		str, err = command.Execute(str)
		if err != nil {
			return "", err
		}
	}
	return str, nil
}

func Z(args []string, r io.Reader, w io.Writer) error {
	config := parseArgs(args)
	if config.Err != nil {
		return config.Err
	}

	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadString('\n')
		isEof := err == io.EOF
		if err != nil && !isEof {
			return err
		}

		var output string
		output, err = config.Execute(line)
		if err != nil {
			return err
		}

		w.Write([]byte(fmt.Sprintln(output)))

		if isEof {
			return nil
		}
	}
}
