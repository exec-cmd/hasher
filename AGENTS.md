## Project overview

`hasher` is a Go 1.26.5 single-binary CLI (`github.com/exec-cmd/hasher`) using
Cobra. The `hash` command hashes one file with `sha256` (default), `sha512`, or
`md5`; output is lowercase hexadecimal.

## Repository layout

- `cmd/hasher/main.go` — application entry point.
- `internal/cmd/` — Cobra commands, flags, argument validation, and output.
- `internal/filehash/` — algorithm selection and file hashing, independent of
  Cobra and terminal output.
- `internal/filehash/hash_test.go` — hashing and algorithm-selection tests.
- `.github/workflows/` — CI verification and tagged cross-platform releases.
- `Makefile` — local formatting, vetting, and build targets.

## Usage and verification

```sh
go run ./cmd/hasher hash <file>
go run ./cmd/hasher hash --alg sha512 <file>
go test ./...
go vet ./...
make build                 # fmt, vet, then build/bin/hasher
```

`--alg` also accepts `sha256` and `md5`. Use `md5` only for compatibility or
checksums, not for security. CI checks formatting, tests, vet, and a trimmed
build; tagged releases build Linux amd64/arm64, Windows amd64, and macOS arm64
binaries.

## Change rules

- Keep CLI code in `internal/cmd` and hashing code in `internal/filehash`.
- Preserve package boundaries and keep changes focused.
- Run `gofmt` (or `make fmt`) after Go changes and add focused tests for new
  behavior.
- Report only checks that were actually run.
