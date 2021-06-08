package main

import (
	"os"
	"strings"

	"github.com/serramatutu/z/internal"
)

func main() {
	args := make([]string, len(os.Args))[:0]
	for _, arg := range os.Args {
		escapedArg := strings.Replace(arg, "\\n", "\n", -1)
		args = append(args, escapedArg)
	}

	err := internal.Z(args, os.Stdin, os.Stdout)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
}
