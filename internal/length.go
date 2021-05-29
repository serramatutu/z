package internal

import (
	_ "embed"
	"fmt"

	"github.com/serramatutu/z/help"
)

type LengthArgs struct{
}

func (LengthArgs) Description() string {
	return help.Help["length"]
}

func Length(in string) (string, error) {
	return fmt.Sprint(len(in)), nil
}
