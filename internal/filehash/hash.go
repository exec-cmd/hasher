package filehash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
)

type Algorithm string

const (
	SHA256 Algorithm = "sha256"
	SHA512 Algorithm = "sha512"
	MD5    Algorithm = "md5"
)

var Algorithms = []Algorithm{
	SHA256,
	SHA512,
	MD5,
}

var ErrInvalidAlgorithm = errors.New("invalid hash algorithm")

func HashFile(path string, h hash.Hash) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func HashFromAlgorithm(a Algorithm) (hash.Hash, error) {
	switch a {
	case SHA256:
		return sha256.New(), nil
	case SHA512:
		return sha512.New(), nil
	case MD5:
		return md5.New(), nil
	default:
		return nil, ErrInvalidAlgorithm
	}
}
