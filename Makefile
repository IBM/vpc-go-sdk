# Makefile

all: build lint tidy

travis-ci: build lint tidy test-unit

build:
	go build ./vpcv1

test-unit:
	go test `go list ./... | grep vpcv1` -v -tags=unit

test-integration:
	go test `go list ./... | grep vpcv1` -v -tags=integration -skipForMockTesting -testCount

lint:
	golangci-lint --timeout=2m run

tidy:
	go mod tidy
