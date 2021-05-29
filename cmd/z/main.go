package main

import (
	"os"

	"github.com/serramatutu/z/internal"
)

func main() {
	err := internal.Z(os.Stdin)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
}
