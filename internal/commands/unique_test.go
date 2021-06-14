package commands_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

var uniqueInput = [][]byte{
	[]byte("a,1,6"),
	[]byte("a,2,6"),
	[]byte("b,1,6"),
	[]byte("c,2,6"),
	[]byte("c,3,6"),
	[]byte("d,3,6"),
}

var uniqueExpectedWholeKey [][]byte = uniqueInput

var uniqueExpectedEmpty = [][]byte{}

var uniqueExpectedIndex0 = [][]byte{
	[]byte("a,1,6"),
	[]byte("b,1,6"),
	[]byte("c,2,6"),
	[]byte("d,3,6"),
}

var uniqueExpectedIndex1 = [][]byte{
	[]byte("a,1,6"),
	[]byte("a,2,6"),
	[]byte("c,3,6"),
}

var uniqueExpectedIndexAllEqual = [][]byte{
	[]byte("a,1,6"),
}

func TestUniqueNilSeparator(t *testing.T) {
	cmd := commands.NewUnique(nil, nil, 0)

	result, err := cmd.Execute(uniqueInput)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with nil separator")
	}

	if !reflect.DeepEqual(result, uniqueExpectedWholeKey) {
		t.Errorf("Wrong result for Unique.Execute with nil separator")
	}
}

func TestUniqueNotFoundSeparator(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(":"), 0)

	result, err := cmd.Execute(uniqueInput)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, uniqueExpectedWholeKey) {
		t.Errorf("Wrong result for Unique.Execute with not found separator")
	}
}

func TestUniqueSeparatorIndex0(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(","), 0)

	result, err := cmd.Execute(uniqueInput)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, uniqueExpectedIndex0) {
		t.Errorf("Wrong result for Unique.Execute with separator")
	}
}

func TestUniqueSeparatorIndex1(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(","), 1)

	result, err := cmd.Execute(uniqueInput)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, uniqueExpectedIndex1) {
		t.Errorf("Wrong result for Unique.Execute with separator")
	}
}

func TestUniqueSeparatorIndexAllEqual(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(","), 2)

	result, err := cmd.Execute(uniqueInput)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, uniqueExpectedIndexAllEqual) {
		t.Errorf("Wrong result for Unique.Execute with separator")
	}
}

func TestUniqueSeparatorIndexOutOfRange(t *testing.T) {
	cmd := commands.NewUnique(nil, regexp.MustCompile(","), 3)

	result, err := cmd.Execute(uniqueInput)
	if err != nil {
		t.Errorf("Unexpected error for Unique.Execute with separator")
	}

	if !reflect.DeepEqual(result, uniqueExpectedIndexAllEqual) {
		t.Errorf("Wrong result for Unique.Execute with separator")
	}
}
