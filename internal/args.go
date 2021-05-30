package internal

import (
	"container/list"
	"errors"
	"fmt"

	"github.com/serramatutu/z/help"
	"github.com/serramatutu/z/internal/commands"
)

func parseCommand(args []string) commands.Command {
	var cmd commands.Command = commands.Invalid{CommandName: args[0]}

	switch args[0] {
	case "length":
		cmd = commands.ParseLength(args[1:])
	case "help":
		cmd = commands.ParseHelp(args[1:])
	}

	return cmd
}

func createError(errText, helpText string) error {
	if errText != "" {
		return fmt.Errorf("error: %s\n\n%s", errText, helpText)
	}
	return errors.New(helpText)
}

// TODO: optimize append
func parseArgs(args []string) Config {
	commandsList := list.New()

	// cannot have underscore at tail
	if len(args) > 1 && args[len(args)-1] == "_" {
		return Config{
			Err: commands.InvalidPipeErr{},
		}
	}

	lastUnderscore := 0
	for i, arg := range args[1:] {
		actualIndex := i + 1

		if arg == "_" {
			// two consecutive underscores
			if lastUnderscore == actualIndex-1 {
				return Config{
					Err: commands.InvalidPipeErr{},
				}
			}

			cmd := parseCommand(args[lastUnderscore+1 : actualIndex])
			if cmd.Err() != nil {
				return Config{
					Err: createError(cmd.Err().Error(), help.Help[cmd.Name()]),
				}
			}

			commandsList.PushBack(cmd)
			lastUnderscore = actualIndex
		}
	}

	if lastUnderscore < len(args) && len(args) > 1 {
		cmd := parseCommand(args[lastUnderscore+1:])
		if cmd.Err() != nil {
			return Config{
				Err: createError(cmd.Err().Error(), help.Help[cmd.HelpFile()]),
			}
		}

		commandsList.PushBack(cmd)
	}

	if commandsList.Len() == 0 {
		return Config{
			Err: createError("no subcommand was given", help.Help["z"]),
		}
	}

	return Config{
		Err:      nil,
		Commands: commandsList,
	}
}
