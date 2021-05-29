package internal

import (
	"bufio"
	"fmt"
	"io"

	"github.com/alexflint/go-arg"
	"github.com/serramatutu/z/help"
)

type ZArgs struct {
	Length *LengthArgs `arg:"subcommand:length"`
}

func (ZArgs) Description() string {
	return help.Help["z"]
}

// TODO: _ separator for piping
func parseArgs() (*ZArgs, error) {
	var args ZArgs

	if err := arg.Parse(&args); err != nil {
		return nil, err
	}

	return &args, nil
}

func execSingleLine(args *ZArgs, line string) (string, error) {
	switch {
	case args.Length != nil:
		return Length(line)
	}

	return "", nil
}

func Z(r io.Reader) (error) {
	args, err := parseArgs()

	if err != nil {
		return err
	}

	reader := bufio.NewReader(r)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF  {
			break
		}

		var output string
		output, err = execSingleLine(args, line)
		if err != nil {
			break
		}

		fmt.Println(output)
	}

	return err
}
