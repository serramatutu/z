package commands

type Command interface {
	Err() error
	Name() string
	HelpFile() string
	Execute([]byte) ([]byte, error)
}
