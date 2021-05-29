package commands

type Command interface {
	Err() error
	Name() string
	Execute(string) (string, error)
}
