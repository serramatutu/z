package commands

type Command interface {
	Err() error
	Name() string
	HelpFile() string
	Execute(string) (string, error)
}
