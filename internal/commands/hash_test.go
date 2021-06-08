package commands_test

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"reflect"
	"testing"

	"github.com/serramatutu/z/internal/commands"
)

func TestHashInvalidAlgorithm(t *testing.T) {
	cmd := commands.NewHash(nil, "invalid")
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	_, err := cmd.Execute(data)

	if err == nil {
		t.Errorf("Unexpected Hash.Execute with invalid algorithm to return error")
	}
}

func TestHashMd5(t *testing.T) {
	cmd := commands.Hash{
		Algorithm: commands.Md5,
	}
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	result, err := cmd.Execute(data)

	if err != nil {
		t.Errorf("Unexpected error for Hash.Execute with MD5 algorithm")
	}

	expected := md5.Sum(data)
	if !reflect.DeepEqual(result, expected[:]) {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestHashSha1(t *testing.T) {
	cmd := commands.NewHash(nil, commands.Sha1)
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	result, err := cmd.Execute(data)

	if err != nil {
		t.Errorf("Unexpected error for Hash.Execute with SHA1 algorithm")
	}

	expected := sha1.Sum(data)
	if !reflect.DeepEqual(result, expected[:]) {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestHashSha224(t *testing.T) {
	cmd := commands.NewHash(nil, commands.Sha224)
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	result, err := cmd.Execute(data)

	if err != nil {
		t.Errorf("Unexpected error for Hash.Execute with SHA1 algorithm")
	}

	expected := sha256.Sum224(data)
	if !reflect.DeepEqual(result, expected[:]) {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestHashSha256(t *testing.T) {
	cmd := commands.NewHash(nil, commands.Sha256)
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	result, err := cmd.Execute(data)

	if err != nil {
		t.Errorf("Unexpected error for Hash.Execute with SHA1 algorithm")
	}

	expected := sha256.Sum256(data)
	if !reflect.DeepEqual(result, expected[:]) {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
