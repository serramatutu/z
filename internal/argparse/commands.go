package argparse

import (
	"regexp"

	"github.com/serramatutu/z/internal/commands"
)

func ParseHelp(args []string) commands.Help {
	var commandName string
	if len(args) > 0 {
		commandName = args[0]
	}

	return commands.Help{
		CommandName: commandName,
	}
}

func ParseJoin(args []string) commands.Join {
	var err error
	var sep []byte

	switch len(args) {
	case 0:
	case 1:
		sep = []byte(args[0])
	default:
		err = ExtraPositionalArgumentErr{
			ArgumentValue: args[1],
		}
		sep = nil
	}

	return commands.NewJoin(err, sep)
}

func ParseLength(args []string) commands.Length {
	var err error
	if len(args) > 0 {
		err = ExtraPositionalArgumentErr{
			ArgumentValue: args[0],
		}
	}

	return commands.NewLength(err)
}

func ParseReplace(args []string) commands.Replace {
	var err error
	var target *regexp.Regexp
	var replacement []byte
	var rangeStart, rangeEnd int

	switch len(args) {
	case 0:
		err = MissingPositionalArgumentsErr{
			ArgumentNames: []string{"pattern", "replace-string"},
		}
	case 1:
		err = MissingPositionalArgumentsErr{
			ArgumentNames: []string{"replace-string"},
		}
	case 3:
		rangeStart, rangeEnd, err = ParseRange(args[2])
		if err != nil {
			err = InvalidPositionalArgumentErr{
				ArgumentName:  "occurrence-range",
				ArgumentValue: args[2],
			}
			break
		}
		fallthrough
	case 2:
		target, err = regexp.Compile(args[0])
		if err != nil {
			err = InvalidPositionalArgumentErr{
				ArgumentName:  "pattern",
				ArgumentValue: args[0],
			}
		}
		replacement = []byte(args[1])
	default:
		err = ExtraPositionalArgumentErr{
			ArgumentValue: args[3],
		}
	}

	return commands.NewReplace(err, target, replacement, rangeStart, rangeEnd)
}

func ParseSplit(args []string) commands.Split {
	var err error
	var sep *regexp.Regexp

	switch len(args) {
	case 0:
	case 1:
		sep, err = regexp.Compile(args[0])
		if err != nil {
			err = InvalidPositionalArgumentErr{
				ArgumentName:  "pattern",
				ArgumentValue: args[0],
			}
		}
	default:
		err = ExtraPositionalArgumentErr{
			ArgumentValue: args[1],
		}
		sep = nil
	}

	return commands.NewSplit(err, sep)
}
