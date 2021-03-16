# Makefile

all: build lint tidy

travis-ci: build lint tidy unittest

build:
	go build ./...

unittest:
	go test ./...

lint:
	golangci-lint run

tidy:
	go mod tidy
