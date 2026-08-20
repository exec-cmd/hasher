.DEFAULT_GOAL := build

.PHONY: fmt vet test build

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	mkdir -p build/bin
	go build -o build/bin/hasher ./cmd/hasher
