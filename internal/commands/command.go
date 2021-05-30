package commands

type Command interface {
	Err() error
	Name() string
	HelpFile() string
}

type MapCommand interface {
	Execute([]byte) ([]byte, error)
}

type SplitCommand interface {
	Execute([]byte) ([][]byte, error)
}

type JoinCommand interface {
	Execute([][]byte) ([]byte, error)
}
