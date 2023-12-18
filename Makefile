

VERSION := $(shell git describe --tags --always --dirty)

generate:
	go generate ./...

build:
	go build -o connectors/pulsor-connectors -ldflags "-X 'github.com/conduitio/conduit-connector-file.version=${VERSION}'" *.go