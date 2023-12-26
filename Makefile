

VERSION := $(shell git describe --tags --always --dirty)

generate:
	go generate ./...

build:
	go build -o connectors/pulsar-connector -ldflags "-X 'github.com/brharrelldev/pulsar-connector.version=${VERSION}'" cmd/main.go