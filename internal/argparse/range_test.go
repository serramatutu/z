package argparse_test

import (
	"testing"

	"github.com/serramatutu/z/internal/argparse"
)

func TestParseRangeInvalid(t *testing.T) {
	_, _, err := argparse.ParseRange("x:y")
	if err == nil {
		t.Errorf("Expected error for ParseRange with invalid arguments")
	}
}

func TestParseRangeFull(t *testing.T) {
	start, end, err := argparse.ParseRange("1:2")
	if err != nil {
		t.Errorf("Unexpected error for ParseRange")
	}

	if start != 1 || end != 2 {
		t.Errorf("Expected 1:2 but got %v:%v for ParseRange", start, end)
	}
}

func TestParseRangeStart(t *testing.T) {
	start, end, err := argparse.ParseRange("1:")
	if err != nil {
		t.Errorf("Unexpected error for ParseRange")
	}

	if start != 1 || end != 0 {
		t.Errorf("Expected 1:0 but got %v:%v for ParseRange", start, end)
	}
}

func TestParseRangeEnd(t *testing.T) {
	start, end, err := argparse.ParseRange(":5")
	if err != nil {
		t.Errorf("Unexpected error for ParseRange")
	}

	if start != 0 || end != 5 {
		t.Errorf("Expected 0:5 but got %v:%v for ParseRange", start, end)
	}
}
