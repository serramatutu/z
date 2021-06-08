package commands

import "fmt"

type Count struct {
	err error
}

func (c Count) Err() error {
	return c.err
}

func (Count) Name() string {
	return "count"
}

func (Count) HelpFile() string {
	return "count"
}

func (c Count) Execute(in [][]byte) ([]byte, error) {
	if c.err != nil {
		return nil, c.err
	}

	return []byte(fmt.Sprint(len(in))), nil
}

func NewCount(err error) Count {
	return Count{
		err: err,
	}
}
