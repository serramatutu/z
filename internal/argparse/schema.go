package argparse

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type argument interface {
	Name() string
	Optional() bool

	Err() error
	Parse(string)
	BecomeDefault()
}

func isWholeRegexMatch(regex *regexp.Regexp, val string) bool {
	matches := regex.FindAllString(val, -1)
	if len(matches) != 1 {
		return false
	}

	match := matches[0]
	return len(match) == len(val)
}

type stringArgument struct {
	name         string
	optional     bool
	defaultValue string
	pattern      string
	err          error
	value        string
}

func (a stringArgument) Optional() bool {
	return a.optional
}

func (a stringArgument) Name() string {
	return a.name
}

func (a stringArgument) Err() error {
	return a.err
}

func (a stringArgument) Value() string {
	return a.value
}

func (a *stringArgument) BecomeDefault() {
	a.value = a.defaultValue
}

func (a *stringArgument) Parse(in string) {
	if a.pattern == "" {
		a.value = in
	} else {
		regex := regexp.MustCompile(a.pattern)
		if !isWholeRegexMatch(regex, in) {
			a.err = InvalidPositionalArgumentErr{
				ArgumentName:  a.name,
				ArgumentValue: in,
			}
		} else {
			a.value = in
		}
	}

}

type patternArgument struct {
	name         string
	optional     bool
	defaultValue *regexp.Regexp
	err          error
	value        *regexp.Regexp
}

func (a patternArgument) Optional() bool {
	return a.optional
}

func (a patternArgument) Name() string {
	return a.name
}

func (a patternArgument) Err() error {
	return a.err
}

func (a patternArgument) Value() *regexp.Regexp {
	return a.value
}

func (a *patternArgument) BecomeDefault() {
	a.value = a.defaultValue
}

func (a *patternArgument) Parse(in string) {
	pattern, err := regexp.Compile(in)
	if err != nil {
		a.err = InvalidPositionalArgumentErr{
			ArgumentName:  a.name,
			ArgumentValue: in,
		}
	} else {
		a.value = pattern
	}
}

type InvalidRangeErr struct {
	Input string
}

func (err InvalidRangeErr) Error() string {
	return fmt.Sprintf("invalid range '%s'", err.Input)
}

type rangeArgument struct {
	name     string
	optional bool
	err      error
	start    int
	end      int
}

func (a rangeArgument) Optional() bool {
	return a.optional
}

func (a rangeArgument) Name() string {
	return a.name
}

func (a rangeArgument) Err() error {
	return a.err
}

func (a rangeArgument) Start() int {
	return a.start
}

func (a rangeArgument) End() int {
	return a.end
}

func (a *rangeArgument) BecomeDefault() {
	a.start = 0
	a.end = 0
}

var rangeExpr *regexp.Regexp = regexp.MustCompile("[0-9]?:(-?[0-9]+)?")

func (a *rangeArgument) Parse(in string) {
	if !isWholeRegexMatch(rangeExpr, in) {
		a.err = InvalidPositionalArgumentErr{
			ArgumentName:  a.name,
			ArgumentValue: in,
		}
		return
	}

	invalidErr := InvalidRangeErr{
		Input: in,
	}

	splitRange := strings.Split(in, ":")
	if len(splitRange) != 2 {
		a.err = invalidErr
		return
	}

	var range64 int64
	if splitRange[0] != "" {
		range64, a.err = strconv.ParseInt(splitRange[0], 10, 0)
		if a.err != nil {
			return
		}
		a.start = int(range64)
	}

	if splitRange[1] != "" {
		range64, a.err = strconv.ParseInt(splitRange[1], 10, 0)
		if a.err != nil {
			return
		}
		a.end = int(range64)
	}
}

func parseSchema(args []string, schema []argument) error {
	switch {
	case len(args) > len(schema):
		return ExtraPositionalArgumentErr{
			ArgumentValue: args[len(args)-1],
		}
	case len(schema) > len(args) && !schema[len(args)].Optional():
		argNames := make([]string, len(schema)-len(args))
		for i, arg := range schema[len(args):] {
			if arg.Optional() {
				argNames = argNames[:i]
				break
			}
			argNames[i] = arg.Name()
		}

		return MissingPositionalArgumentsErr{
			ArgumentNames: argNames,
		}
	}

	for i, argValue := range args {
		schema[i].Parse(argValue)
		err := schema[i].Err()
		if err != nil {
			return err
		}
	}

	for _, optionalArg := range schema[len(args):] {
		optionalArg.BecomeDefault()
	}

	return nil
}
