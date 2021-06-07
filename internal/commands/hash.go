package commands

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
)

type HashAlgorithm string

const (
	Md5    = HashAlgorithm("md5")
	Sha1   = HashAlgorithm("sha1")
	Sha224 = HashAlgorithm("sha224")
	Sha256 = HashAlgorithm("sha256")
)

func (a HashAlgorithm) Validate() error {
	switch a {
	case Md5, Sha1, Sha224, Sha256:
		return nil
	}
	return errors.New("Invalid hash algorithm")
}

type Hash struct {
	err       error
	Algorithm HashAlgorithm
}

func (h Hash) Err() error {
	return h.err
}

func (Hash) Name() string {
	return "hash"
}

func (Hash) HelpFile() string {
	return "hash"
}

type InvalidAlgorithmErr struct {
	Algorithm string
}

func (err InvalidAlgorithmErr) Error() string {
	return fmt.Sprintf("Invalid provided algorithm \"%s\"", err.Algorithm)
}

func (h Hash) Execute(in []byte) ([]byte, error) {
	switch h.Algorithm {
	case Md5:
		sum := md5.Sum(in)
		return sum[:], nil
	case Sha1:
		sum := sha1.Sum(in)
		return sum[:], nil
	case Sha224:
		sum := sha256.Sum224(in)
		return sum[:], nil
	case Sha256:
		sum := sha256.Sum256(in)
		return sum[:], nil
	}

	return nil, InvalidAlgorithmErr{
		Algorithm: string(h.Algorithm),
	}
}

func NewHash(err error, algorithm HashAlgorithm) Hash {
	return Hash{
		err:       err,
		Algorithm: algorithm,
	}
}
