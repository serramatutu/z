package commands

import "testing"

func TestLengthExecute(t *testing.T) {
	l := Length{}
	result, err := l.Execute("1234")
	if err != nil {
		t.Errorf("Length.Execute should never return error")
	}

	if result != "4" {
		t.Errorf("Length.Execute result is wrong")
	}
}

func TestParseLengthNoArgs(t *testing.T) {
	args := []string{}
	length := ParseLength(args)

	if length.Err() != nil {
		t.Errorf("ParseLength should not return error when no args are given")
	}
}

func TestParseLengthWithArgs(t *testing.T) {
	args := []string{"arg"}
	length := ParseLength(args)

	if length.Err() == nil {
		t.Errorf("ParseLength should return error when args are given")
	}
}
