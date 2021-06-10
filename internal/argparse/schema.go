package argparse

import (
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

type numberArgument struct {
	name         string
	optional     bool
	defaultValue int
	err          error
	value        int
}

func (a numberArgument) Optional() bool {
	return a.optional
}

func (a numberArgument) Name() string {
	return a.name
}

func (a numberArgument) Err() error {
	return a.err
}

func (a numberArgument) Value() int {
	return a.value
}

func (a *numberArgument) BecomeDefault() {
	a.value = a.defaultValue
}

var numberPattern *regexp.Regexp = regexp.MustCompile("^[0-9]+$")

func (a *numberArgument) Parse(in string) {
	if !numberPattern.MatchString(in) {
		a.err = InvalidPositionalArgumentErr{
			ArgumentName:  a.name,
			ArgumentValue: in,
		}
	} else {
		val, _ := strconv.ParseInt(in, 10, 8)
		a.value = int(val)
	}
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
		if !regex.MatchString(in) {
			a.err = InvalidPositionalArgumentErr{
				ArgumentName:  a.name,
				ArgumentValue: in,
			}
		} else {
			a.value = in
		}
	}
}

type enumArgument struct {
	name         string
	optional     bool
	defaultValue string
	err          error
	options      []string
	value        string
}

func (a enumArgument) Optional() bool {
	return a.optional
}

func (a enumArgument) Name() string {
	return a.name
}

func (a enumArgument) Err() error {
	return a.err
}

func (a enumArgument) Value() string {
	return a.value
}

func (a *enumArgument) BecomeDefault() {
	a.value = a.defaultValue
}

func (a *enumArgument) Parse(in string) {
	for _, val := range a.options {
		if in == val {
			a.value = val
			return
		}
	}

	a.err = InvalidPositionalArgumentErr{
		ArgumentName:  a.name,
		ArgumentValue: in,
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

var rangeExpr *regexp.Regexp = regexp.MustCompile("^[0-9]?:(-?[0-9]+)?$")

func (a *rangeArgument) Parse(in string) {
	if !rangeExpr.MatchString(in) {
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
