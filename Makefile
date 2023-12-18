

VERSION := $(shell git describe --tags --always --dirty)

generate:
	go generate ./...

build:
	go build -o connectors/pulsor-connectors -ldflags "-X 'github.com/brharrelldev/pulsar-connector.version=${VERSION}'" *.go