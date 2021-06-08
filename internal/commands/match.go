package commands

import (
	"errors"
	"regexp"
)

type Match struct {
	err error

	Pattern *regexp.Regexp
}

func (m Match) Err() error {
	return m.err
}

func (Match) Name() string {
	return "match"
}

func (Match) HelpFile() string {
	return "match"
}

func (m Match) Execute(in []byte) ([][]byte, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.Pattern.FindAll(in, -1), nil
}

func NewMatch(err error, pattern *regexp.Regexp) Match {
	if err == nil && pattern == nil {
		err = errors.New("cannot execute Match with nil pattern")
	}

	return Match{
		err:     err,
		Pattern: pattern,
	}
}
