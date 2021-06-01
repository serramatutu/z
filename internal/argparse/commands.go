package argparse

import (
	"github.com/serramatutu/z/internal/commands"
)

func ParseHelp(args []string) commands.Help {
	commandName := stringArgument{
		name:     "command-name",
		optional: true,
	}
	schema := []argument{
		&commandName,
	}
	err := parseArgsSchema(args, schema)
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
	err := parseArgsSchema(args, schema)

	var sep []byte
	if separator.Value() != "" {
		sep = []byte(separator.Value())
	}
	return commands.NewJoin(err, sep)
}

func ParseLength(args []string) commands.Length {
	schema := []argument{}
	err := parseArgsSchema(args, schema)
	return commands.NewLength(err)
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
	err := parseArgsSchema(args, schema)

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
		defaultValue: nil,
	}
	schema := []argument{
		&pattern,
	}
	err := parseArgsSchema(args, schema)

	return commands.NewSplit(err, pattern.Value())
}
