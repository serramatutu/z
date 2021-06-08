package argparse

import (
	"regexp"

	"github.com/serramatutu/z/internal/commands"
)

func ParseCount(args []string) commands.Count {
	schema := []argument{}
	err := parseSchema(args, schema)
	return commands.NewCount(err)
}

func ParseHash(args []string) commands.Hash {
	algorithm := enumArgument{
		name:     "algorithm",
		optional: false,
		options: []string{
			"md5",
			"sha1",
			"sha224",
			"sha256",
		},
	}
	schema := []argument{
		&algorithm,
	}
	err := parseSchema(args, schema)
	return commands.NewHash(err, commands.HashAlgorithm(algorithm.Value()))
}

func ParseHelp(args []string) commands.Help {
	commandName := stringArgument{
		name:     "command-name",
		optional: true,
	}
	schema := []argument{
		&commandName,
	}
	err := parseSchema(args, schema)
	if err != nil {
		return commands.Help{
			CommandName: "z",
		}
	}

	return commands.Help{
		CommandName: commandName.Value(),
	}
}

func ParseJoin(args []string) commands.Join {
	separator := stringArgument{
		name:         "separator",
		optional:     true,
		defaultValue: "",
	}
	schema := []argument{
		&separator,
	}
	err := parseSchema(args, schema)

	var sep []byte
	if err == nil {
		sep = []byte(separator.Value())
	}
	return commands.NewJoin(err, sep)
}

func ParseLength(args []string) commands.Length {
	mode := enumArgument{
		name:         "mode",
		optional:     true,
		defaultValue: "bytes",
		options: []string{
			"unicode",
			"bytes",
		},
	}
	schema := []argument{
		&mode,
	}
	err := parseSchema(args, schema)
	return commands.NewLength(err, commands.LengthMode(mode.Value()))
}

func ParseMatch(args []string) commands.Match {
	pattern := patternArgument{
		name:     "pattern",
		optional: false,
	}
	schema := []argument{
		&pattern,
	}
	err := parseSchema(args, schema)

	return commands.NewMatch(err, pattern.Value())
}

func ParseReplace(args []string) commands.Replace {
	pattern := patternArgument{
		name:     "pattern",
		optional: false,
	}
	replacement := stringArgument{
		name:     "replace-string",
		optional: false,
	}
	rangeArg := rangeArgument{
		name:     "occurrence-range",
		optional: true,
	}
	schema := []argument{
		&pattern,
		&replacement,
		&rangeArg,
	}
	err := parseSchema(args, schema)

	return commands.NewReplace(
		err,
		pattern.Value(),
		[]byte(replacement.Value()),
		rangeArg.Start(),
		rangeArg.End(),
	)
}

func ParseSplit(args []string) commands.Split {
	pattern := patternArgument{
		name:         "pattern",
		optional:     true,
		defaultValue: regexp.MustCompile("\n"),
	}
	schema := []argument{
		&pattern,
	}
	err := parseSchema(args, schema)

	return commands.NewSplit(err, pattern.Value())
}

func ParseVersion(args []string) commands.Version {
	return commands.Version{}
}
