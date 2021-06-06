package commands

import (
	"fmt"

	"github.com/serramatutu/z/internal/config"
)

type Version struct {
}

func (Version) Err() error {
	return nil
}

func (Version) Name() string {
	return "version"
}

func (Version) HelpFile() string {
	return "z"
}

func (Version) Execute() []byte {
	return []byte(fmt.Sprintf(
		"z %s\n  built at: %s\n  built by: %s\n  based on commit: %s\n  based on repository: %s\n",
		config.Version,
		config.Date,
		config.BuiltBy,
		config.Commit,
		config.Repository,
	))
}
