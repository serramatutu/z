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

func NewConfig(err error, commands *list.List) Config {
	return Config{
		Err:      err,
		Commands: commands,
	}
}

func executeMap(bytes []byte, start *list.Element) ([]byte, *list.Element, error) {
	var lastRan *list.Element
	var err error

	for e := start; e != nil; e = e.Next() {
		switch e.Value.(type) {
		case commands.MapCommand:
			command := e.Value.(commands.MapCommand)
			bytes, err = command.Execute(bytes)
			if err != nil {
				return nil, e, err
			}
		default:
			return bytes, lastRan, nil
		}

		lastRan = e
	}

	return bytes, lastRan, nil
}

var implicitJoin commands.JoinCommand = commands.Join{
	Separator: []byte(""),
}

func executeSplit(bytes []byte, start *list.Element) ([]byte, *list.Element, error) {
	var lastRan *list.Element

	command := start.Value.(commands.SplitCommand)
	splitBytes, err := command.Execute(bytes)
	if err != nil {
		return nil, start, err
	}

	for e := start.Next(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case commands.JoinCommand:
			command := e.Value.(commands.JoinCommand)
			bytes, err = command.Execute(splitBytes)
			if err != nil {
				return nil, e, err
			}
			return bytes, e, nil

		case commands.SplitCommand:
			for i := 0; i < len(splitBytes); i++ {
				splitBytes[i], lastRan, err = executeSplit(splitBytes[i], e)
				if err != nil {
					return nil, e, err
				}
			}
			e = lastRan

		case commands.MapCommand:
			for i := 0; i < len(splitBytes); i++ {
				splitBytes[i], lastRan, err = executeMap(splitBytes[i], e)
				if err != nil {
					return nil, e, err
				}
			}
			e = lastRan
		}
	}

	// TODO: disallow implicit joins
	// return nil, lastRan, commands.ExtraSplitErr{}

	// in this case, return the last _known_ command that was ran
	// (implicit joins are unknown to the config)
	bytes, err = implicitJoin.Execute(splitBytes)
	if err != nil {
		return nil, lastRan, err
	}
	return bytes, lastRan, nil
}

func (config Config) Execute(bytes []byte) ([]byte, error) {
	var err error

	for e := config.Commands.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case commands.JoinCommand:
			return nil, commands.ExtraJoinErr{}
		case commands.MapCommand:
			bytes, e, err = executeMap(bytes, e)
			if err != nil {
				return nil, err
			}
		case commands.SplitCommand:
			bytes, e, err = executeSplit(bytes, e)
			if err != nil {
				return nil, err
			}
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
