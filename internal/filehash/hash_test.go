package filehash_test

import (
	"crypto/sha256"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"reflect"

	"github.com/exec-cmd/hasher/internal/filehash"
)

const (
	textTempFile = "vibe"
	pathTempFile = "test.txt"
	sha256Hex    = "8368ec5f1a153787469953cac4a39bce1c1e9273175a0a39db0810d0c7d0273d"
	sha512Hex    = "15424a2bbc50a595c84797561f72f64dc1842f7b86d8e91976e0017999aa4eea8bb98f6984511dc105e809eac5be812e2aaaf05bdbcf447137610c81a46eb17e"
	md5Hex       = "1876a682f240c24b40d879acfee19fb1"
)

func TestHashFileSHA256(t *testing.T) {
	file := filepath.Join(t.TempDir(), pathTempFile)

	err := os.WriteFile(file, []byte(textTempFile), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	h, err := filehash.HashFromAlgorithm(filehash.SHA256)
	if err != nil {
		t.Fatal(err)
	}

	got, err := filehash.HashFile(file, h)
	if err != nil {
		t.Fatal(err)
	}

	want := sha256Hex

	if got != want {
		t.Errorf("Hash() = %v, want %v", got, want)
	}
}

func TestHashFileSHA512(t *testing.T) {
	file := filepath.Join(t.TempDir(), pathTempFile)

	err := os.WriteFile(file, []byte(textTempFile), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	h, err := filehash.HashFromAlgorithm(filehash.SHA512)
	if err != nil {
		t.Fatal(err)
	}

	got, err := filehash.HashFile(file, h)
	if err != nil {
		t.Fatal(err)
	}

	want := sha512Hex

	if got != want {
		t.Errorf("Hash() = %v, want %v", got, want)
	}
}

func TestHashFileMD5(t *testing.T) {
	file := filepath.Join(t.TempDir(), pathTempFile)

	err := os.WriteFile(file, []byte(textTempFile), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	h, err := filehash.HashFromAlgorithm(filehash.MD5)
	if err != nil {
		t.Fatal(err)
	}

	got, err := filehash.HashFile(file, h)
	if err != nil {
		t.Fatal(err)
	}

	want := md5Hex

	if got != want {
		t.Errorf("Hash() = %v, want %v", got, want)
	}
}

func TestHashFromAlgorithm(t *testing.T) {
	got, err := filehash.HashFromAlgorithm(filehash.Algorithm("sha256"))
	if err != nil {
		t.Fatal(err)
	}

	want := sha256.New()

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("got type %T, want %T", got, want)
	}
}

func TestHashFileInvalidAlgorithm(t *testing.T) {
	_, err := filehash.HashFromAlgorithm(filehash.Algorithm("invalid"))

	if !errors.Is(err, filehash.ErrInvalidAlgorithm) {
		t.Errorf("got error %v, want ErrInvalidAlgorithm", err)
	}
}
