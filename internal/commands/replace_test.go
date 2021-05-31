package commands_test

import (
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func execReplace(data, target, repl string, start, end int) (string, error) {
	cmd := commands.Replace{
		Target:      regexp.MustCompile(target),
		Replacement: []byte(repl),
		RangeStart:  start,
		RangeEnd:    end,
	}
	result, err := cmd.Execute([]byte(data))
	return string(result), err
}

func TestNoTarget(t *testing.T) {
	cmd := commands.Replace{}
	_, err := cmd.Execute([]byte(""))
	if err == nil {
		t.Errorf("Expected Replace to return error when Target is nil")
	}
}

func TestNoMatch(t *testing.T) {
	result, err := execReplace("aaa:bbb:ccc:ddd", "-", "_", 0, 0)
	if err != nil {
		t.Errorf("Unexpected error for Replace.Execute")
	}

	expected := "aaa:bbb:ccc:ddd"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestReplaceByEmpty(t *testing.T) {
	result, err := execReplace("aaa:bbb:ccc:ddd", ":", "", 0, 0)
	if err != nil {
		t.Errorf("Unexpected error for Replace.Execute")
	}

	expected := "aaabbbcccddd"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestFullRange(t *testing.T) {
	result, err := execReplace("aaa:bbb:ccc:ddd", ":", "-", 0, 0)
	if err != nil {
		t.Errorf("Unexpected error for Replace.Execute")
	}

	expected := "aaa-bbb-ccc-ddd"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestEmptyRange(t *testing.T) {
	result, err := execReplace("aaa:bbb:ccc:ddd", ":", "-", 2, 1)
	if err != nil {
		t.Errorf("Unexpected error for Replace.Execute")
	}

	expected := "aaa:bbb:ccc:ddd"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestRangeBeginning(t *testing.T) {
	result, err := execReplace("aaa:bbb:ccc:ddd", ":", "-", 1, 0)
	if err != nil {
		t.Errorf("Unexpected error for Replace.Execute")
	}

	expected := "aaa:bbb-ccc-ddd"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestRangeEndPositiveLessMatches(t *testing.T) {
	result, err := execReplace("aaa:bbb:ccc:ddd", ":", "-", 1, 2)
	if err != nil {
		t.Errorf("Unexpected error for Replace.Execute")
	}

	expected := "aaa:bbb-ccc:ddd"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestRangeEndPositiveMoreMatches(t *testing.T) {
	result, err := execReplace("aaa:bbb:ccc:ddd", ":", "-", 1, 6)
	if err != nil {
		t.Errorf("Unexpected error for Replace.Execute")
	}

	expected := "aaa:bbb-ccc-ddd"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestRangeEndNegative(t *testing.T) {
	result, err := execReplace("aaa:bbb:ccc:ddd", ":", "-", 1, -1)
	if err != nil {
		t.Errorf("Unexpected error for Replace.Execute")
	}

	expected := "aaa:bbb-ccc:ddd"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
