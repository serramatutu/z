package internal

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteLength(t *testing.T) {
	args := []string{"z", "length"}
	in := strings.NewReader("1234\n\n")

	var out bytes.Buffer
	Z(args, in, &out)

	expected := []byte("6")

	if !bytes.Equal(out.Bytes(), []byte("6")) {
		t.Errorf("Expected '%s' as Z output but got '%s'", expected, out.String())
	}
}
