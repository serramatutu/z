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
func parseArgs() *ZArgs {
	var args ZArgs

	p := arg.MustParse(&args)
	if p.Subcommand() == nil {
		p.Fail(help.Help["z"])
	}

	return &args
}

func execSingleLine(args *ZArgs, line string) (string, error) {
	switch {
	case args.Length != nil:
		return Length(line)
	}

	return "", nil
}

func Z(r io.Reader) error {
	args := parseArgs()

	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadString('\n')
		isEof := err == io.EOF
		if err != nil && !isEof {
			return err
		}

		var output string
		output, err = execSingleLine(args, line)
		if err != nil {
			return err
		}

		fmt.Println(output)

		if isEof {
			return nil
		}
	}
}
