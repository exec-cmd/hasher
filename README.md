# hasher

Russian version: [README_RU.md](README_RU.md)

## About

`hasher` is a small Go command-line utility for calculating file hashes.

Supported algorithms:

- `sha256` - used by default;
- `sha512`;
- `md5` - for compatibility and checksums, not for data protection.

The result is printed as a lowercase hexadecimal string.

## Installation

Download the appropriate executable from the [Releases](https://github.com/exec-cmd/hasher/releases) page and place it in a directory included in your `PATH`.

On Linux and macOS, you may need to make the file executable after downloading it:

```sh
chmod +x hasher
```

### Building from source

Go `1.26.5` or newer is required to build the project.

Clone the repository and build the binary:

```sh
git clone https://github.com/exec-cmd/hasher.git
cd hasher
go build -o hasher ./cmd/hasher
```

The binary will be available in the current directory after the build.

## Usage

Hash a file with `sha256`:

```sh
./hasher hash path/to/file
```

Select an algorithm with `--alg` or its short form, `-a`:

```sh
./hasher hash --alg sha512 path/to/file
./hasher hash -a md5 path/to/file
```

Display help:

```sh
./hasher --help
./hasher hash --help
```
