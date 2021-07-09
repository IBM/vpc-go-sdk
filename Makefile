# Makefile

all: build lint tidy

travis-ci: build lint tidy unittest

build:
	go build ./...

unittest:
	go test ./...

alltest:
	go test `go list ./... | grep vpcv1` -v -tags=integration -skipForMockTesting -testCount

lint:
	golangci-lint --timeout=2m run

tidy:
	go mod tidy
