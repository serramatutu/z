package argparse

import (
	"reflect"
	"regexp"
	"testing"
)

func TestParseStringArgumentNoPattern(t *testing.T) {
	arg := stringArgument{
		name:     "arg-name",
		optional: false,
	}
	arg.Parse("value")

	if arg.Err() != nil {
		t.Errorf("stringArgument.Parse() with no pattern should not return error")
	}

	if arg.Value() != "value" {
		t.Errorf("Expected stringArgument without pattern to assume any value")
	}
}

func TestParseStringArgumentWithPatternOk(t *testing.T) {
	arg := stringArgument{
		name:     "arg-name",
		optional: false,
		pattern:  "[A-z]+",
	}
	arg.Parse("Value")

	if arg.Err() != nil {
		t.Errorf("stringArgument.Parse() with matching pattern should not return error")
	}

	if arg.Value() != "Value" {
		t.Errorf("Expected stringArgument with matching pattern to assume given value")
	}
}

func TestParseStringArgumentWithPatternNotWholeMatch(t *testing.T) {
	arg := stringArgument{
		name:     "arg-name",
		optional: false,
		pattern:  "[A-z]+",
	}
	arg.Parse("Value1234")

	if arg.Err() == nil {
		t.Errorf("stringArgument.Parse() with partially matching pattern should return error")
	}
}

func TestParsePatternArgumentOk(t *testing.T) {
	arg := patternArgument{
		name:     "arg-name",
		optional: false,
	}
	arg.Parse("[A-z]+")

	if arg.Err() != nil {
		t.Errorf("patternArgument.Parse() with compilable pattern should not return error")
	}

	expected := regexp.MustCompile("[A-z]+")
	if !reflect.DeepEqual(*arg.Value(), *expected) {
		t.Errorf("Expected patternArgument with compilable pattern to assume compiled regex value")
	}
}

func TestParsePatternArgumentInvalidPattern(t *testing.T) {
	arg := patternArgument{
		name:     "arg-name",
		optional: false,
	}
	arg.Parse("[A-z][")

	if arg.Err() == nil {
		t.Errorf("patternArgument.Parse() with invalid pattern should return error")
	}
}

func TestRangeArgumentInvalid(t *testing.T) {
	arg := rangeArgument{
		name:     "arg-name",
		optional: false,
	}
	arg.Parse("x:y")

	if arg.Err() == nil {
		t.Errorf("Expected error for rangeArgument.Parse() with invalid arguments")
	}
}

func TestParseRangeFull(t *testing.T) {
	arg := rangeArgument{
		name:     "arg-name",
		optional: false,
	}
	arg.Parse("1:2")
	if arg.Err() != nil {
		t.Errorf("Unexpected error for ParseRange")
	}

	if arg.Start() != 1 || arg.End() != 2 {
		t.Errorf("Expected 1:2 but got %v:%v for ParseRange", arg.Start(), arg.End())
	}
}

func TestParseRangeStart(t *testing.T) {
	arg := rangeArgument{
		name:     "arg-name",
		optional: false,
	}
	arg.Parse("1:")
	if arg.Err() != nil {
		t.Errorf("Unexpected error for ParseRange")
	}

	if arg.Start() != 1 || arg.End() != 0 {
		t.Errorf("Expected 1:0 but got %v:%v for ParseRange", arg.Start(), arg.End())
	}
}

func TestParseRangeEnd(t *testing.T) {
	arg := rangeArgument{
		name:     "arg-name",
		optional: false,
	}
	arg.Parse(":5")
	if arg.Err() != nil {
		t.Errorf("Unexpected error for ParseRange")
	}

	if arg.Start() != 0 || arg.End() != 5 {
		t.Errorf("Expected 0:5 but got %v:%v for ParseRange", arg.Start(), arg.End())
	}
}
