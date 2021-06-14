package utility_test

import (
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/utility"
)

var nilSep *regexp.Regexp = nil
var commaSep *regexp.Regexp = regexp.MustCompile(",")
var tabSep *regexp.Regexp = regexp.MustCompile("\t")
var commaIn string = "a,b,c,d"

func TestKeyNoSeparator(t *testing.T) {
	result := utility.ExtractKeyFromFields([]byte(commaIn), nilSep, 0)

	if result != commaIn {
		t.Errorf("Expected '%s' as key but got '%s' for nil separator", commaIn, result)
	}
}

func TestKeyNoMatchIndex0(t *testing.T) {
	result := utility.ExtractKeyFromFields([]byte(commaIn), tabSep, 0)

	if result != commaIn {
		t.Errorf("Expected '%s' as key but got '%s' for no matches with index 0", commaIn, result)
	}
}

func TestKeyNoMatchIndexGreaterThan0(t *testing.T) {
	result := utility.ExtractKeyFromFields([]byte(commaIn), tabSep, 1)

	if result != "" {
		t.Errorf("Expected '%s' as key but got '%s' for no matches with index > 0", "", result)
	}
}

func TestKeyMatchIndex0(t *testing.T) {
	result := utility.ExtractKeyFromFields([]byte(commaIn), commaSep, 0)

	if result != "a" {
		t.Errorf("Expected '%s' as key but got '%s' for valid match with index 0", "", result)
	}
}

func TestKeyMatchIndex1(t *testing.T) {
	result := utility.ExtractKeyFromFields([]byte(commaIn), commaSep, 1)

	if result != "b" {
		t.Errorf("Expected '%s' as key but got '%s' for valid match with index 1", "b", result)
	}
}

func TestKeyMatchLastIndex(t *testing.T) {
	result := utility.ExtractKeyFromFields([]byte(commaIn), commaSep, 3)

	if result != "d" {
		t.Errorf("Expected '%s' as key but got '%s' for valid match with last index", "d", result)
	}
}

func TestKeyMatchIndexOutOfRange(t *testing.T) {
	result := utility.ExtractKeyFromFields([]byte(commaIn), commaSep, 99)

	if result != "" {
		t.Errorf("Expected '%s' as key but got '%s' for valid match with index out of range", "", result)
	}
}
