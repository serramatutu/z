package commands_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

var input = [][]byte{
	[]byte("a,1,6"),
	[]byte("a,2,6"),
	[]byte("b,1,6"),
	[]byte("c,2,6"),
	[]byte("c,3,6"),
	[]byte("d,3,6"),
}

var expectedWholeKey [][]byte = input

var expectedEmpty = [][]byte{}

var expectedIndex0 = [][]byte{
	[]byte("a,1,6"),
	[]byte("b,1,6"),
	[]byte("c,2,6"),
	[]byte("d,3,6"),
}

var expectedIndex1 = [][]byte{
	[]byte("a,1,6"),
	[]byte("a,2,6"),
	[]byte("c,3,6"),
}

var expectedIndexAllEqual = [][]byte{
	[]byte("a,1,6"),
}

func TestUniqueNilSeparator(t *testing.T) {
	cmd := commands.NewUnique(nil, nil, 0)

	result, err := cmd.Execute(input)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with nil separator")
	}

	if !reflect.DeepEqual(result, expectedWholeKey) {
		t.Errorf("Wrong result for Unique.Execute with nil separator")
	}
}

func TestUniqueNotFoundSeparator(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(":"), 0)

	result, err := cmd.Execute(input)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, expectedWholeKey) {
		t.Errorf("Wrong result for Unique.Execute with not found separator")
	}
}

func TestUniqueSeparatorIndex0(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(","), 0)

	result, err := cmd.Execute(input)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, expectedIndex0) {
		t.Errorf("Wrong result for Unique.Execute with separator")
	}
}

func TestUniqueSeparatorIndex1(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(","), 1)

	result, err := cmd.Execute(input)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, expectedIndex1) {
		t.Errorf("Wrong result for Unique.Execute with separator")
	}
}

func TestUniqueSeparatorIndexAllEqual(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(","), 2)

	result, err := cmd.Execute(input)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, expectedIndexAllEqual) {
		t.Errorf("Wrong result for Unique.Execute with separator")
	}
}

func TestUniqueSeparatorIndexOutOfRange(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(","), 3)

	result, err := cmd.Execute(input)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, expectedIndexAllEqual) {
		t.Errorf("Wrong result for Unique.Execute with separator")
	}
}
