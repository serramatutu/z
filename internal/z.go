package internal

import (
	"container/list"
	"io"
	"io/ioutil"

	"github.com/serramatutu/z/internal/argparse"
	"github.com/serramatutu/z/internal/commands"
	"github.com/serramatutu/z/internal/config"
)

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
	command := start.Value.(commands.SplitCommand)
	splitBytes, err := command.Execute(bytes)
	if err != nil {
		return nil, start, err
	}

	lastRan := start

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

		case commands.ArrayMapCommand:
			command := e.Value.(commands.ArrayMapCommand)
			splitBytes, err = command.Execute(splitBytes)
			if err != nil {
				return nil, e, err
			}

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

func executeConfig(c config.Config, bytes []byte) ([]byte, error) {
	var err error

	for e := c.Commands.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case commands.JoinCommand:
			return nil, argparse.ExtraJoinErr{}
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

func Z(zArgs []string, r io.Reader, w io.Writer) error {
	c := argparse.ParseArgs(zArgs)
	if c.Err != nil {
		return c.Err
	}

	// "help" and "version" commands need to exit before reading input
	switch c.Commands.Front().Value.(type) {
	case commands.SingleExecCommand:
		cmd := c.Commands.Front().Value.(commands.SingleExecCommand)
		w.Write(cmd.Execute())
		return nil
	}

	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var output []byte
	output, err = executeConfig(c, contents)
	if err != nil {
		return err
	}

	w.Write(output)

	return nil
}
