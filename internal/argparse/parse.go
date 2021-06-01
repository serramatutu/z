package argparse

import (
	"container/list"
	"fmt"

	"github.com/serramatutu/z/internal/commands"
	"github.com/serramatutu/z/internal/config"
)

type ArgumentErr struct {
	CommandName string
	ErrText     string
}

func (err ArgumentErr) Error() string {
	switch err.CommandName {
	case "help":
		return err.ErrText
	default:
		runHelpText := fmt.Sprintf("run \"z help %s\" to learn more.\n", err.CommandName)
		if err.ErrText != "" {
			return fmt.Sprintf("error: %s\n%s", err.ErrText, runHelpText)
		}
		return runHelpText
	}
}

func parseCommand(args []string) commands.Command {
	var cmd commands.Command = commands.Invalid{CommandName: args[0]}

	switch args[0] {
	case "help":
		cmd = ParseHelp(args[1:])
	case "join":
		cmd = ParseJoin(args[1:])
	case "length":
		cmd = ParseLength(args[1:])
	case "split":
		cmd = ParseSplit(args[1:])
	case "replace":
		cmd = ParseReplace(args[1:])
	}

	return cmd
}

func ParseArgs(args []string) config.Config {
	commandsList := list.New()

	// cannot have underscore at tail
	if len(args) > 1 && args[len(args)-1] == "_" {
		return config.Config{
			Err: InvalidPipeErr{},
		}
	}

	lastUnderscore := 0
	for i, arg := range args[1:] {
		actualIndex := i + 1

		if arg == "_" {
			// two consecutive underscores
			if lastUnderscore == actualIndex-1 {
				return config.Config{
					Err: InvalidPipeErr{},
				}
			}

			cmd := parseCommand(args[lastUnderscore+1 : actualIndex])
			if cmd.Err() != nil {
				return config.Config{
					Err: ArgumentErr{
						ErrText:     cmd.Err().Error(),
						CommandName: cmd.Name(),
					},
				}
			}

			commandsList.PushBack(cmd)
			lastUnderscore = actualIndex
		}
	}

	if lastUnderscore < len(args) && len(args) > 1 {
		cmd := parseCommand(args[lastUnderscore+1:])
		if cmd.Err() != nil {
			return config.Config{
				Err: ArgumentErr{
					ErrText:     cmd.Err().Error(),
					CommandName: cmd.Name(),
				},
			}
		}

		commandsList.PushBack(cmd)
	}

	if commandsList.Len() == 0 {
		return config.Config{
			Err: ArgumentErr{
				ErrText:     "no subcommand was given",
				CommandName: "z",
			},
		}
	}

	return config.Config{
		Err:      nil,
		Commands: commandsList,
	}
}