package config

import "container/list"

type Config struct {
	Err      error
	Commands *list.List
}

func NewConfig(err error, commands *list.List) Config {
	return Config{
		Err:      err,
		Commands: commands,
	}
}
