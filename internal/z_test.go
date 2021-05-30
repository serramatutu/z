package internal

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteLength(t *testing.T) {
	args := []string{"z", "length"}
	in := strings.NewReader("1234\n12345\n123456\n1234567")

	var out bytes.Buffer
	Z(args, in, &out)

	expected := "5\n6\n7\n7"

	if out.String() != expected {
		t.Errorf("Expected '%s' as Z output but got '%s'", expected, out.String())
	}
}
