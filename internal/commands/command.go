package commands

type Command interface {
	Err() error
	Name() string
	HelpFile() string
}

type SingleExecCommand interface {
	Execute() []byte
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
