# Makefile

all: build lint tidy

travis-ci: build lint alltestgen1 alltestgen2 tidy

build:
	go build ./...

alltestgen1:
	cd vpcclassicv1 && go test -run TestVpcClassicV1 && go test `go list ./...` -v -tags=integration -skipForMockTesting -testCount && cd ..

alltestgen2:
	cd vpcv1 && go test -run TestVpcV1 && go test `go list ./...` -v -tags=integration -skipForMockTesting -testCount && cd ..

lint:
	golangci-lint run

tidy:
	go mod tidy
