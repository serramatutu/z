package internal

import (
	"container/list"
	"errors"
	"fmt"
	"os"

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
func parseArgs() Config {
	commands := list.New()
	lastUnderscore := 0
	for i, arg := range os.Args {
		if arg == "_" {
			cmd := parseCommand(os.Args[lastUnderscore+1 : i])
			if cmd.Err() != nil {
				return Config{
					Err:      createError(cmd.Err().Error(), help.Help[cmd.Name()]),
					Commands: nil,
				}
			}

			commands.PushBack(cmd)
			lastUnderscore = i
		}
	}

	if lastUnderscore < len(os.Args) && len(os.Args) > 1 {
		cmd := parseCommand(os.Args[lastUnderscore+1 : len(os.Args)])
		if cmd.Err() != nil {
			return Config{
				Err:      createError(cmd.Err().Error(), help.Help[cmd.HelpFile()]),
				Commands: nil,
			}
		}

		commands.PushBack(cmd)
	}

	if commands.Len() == 0 {
		return Config{
			Err:      createError("no subcommand was given", help.Help["z"]),
			Commands: nil,
		}
	}

	return Config{
		Err:      nil,
		Commands: commands,
	}
}
