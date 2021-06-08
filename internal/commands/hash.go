package commands

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
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

func (h Hash) Execute(in []byte) ([]byte, error) {
	if h.err != nil {
		return nil, h.err
	}

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

	// should never reach this
	return nil, nil
}

func NewHash(err error, algorithm HashAlgorithm) Hash {
	if err == nil {
		err = algorithm.Validate()
	}

	return Hash{
		err:       err,
		Algorithm: algorithm,
	}
}
