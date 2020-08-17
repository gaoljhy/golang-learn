# note: call scripts from /scripts
export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on

# default exectude all 
all: fmt build

build: main

# compile assets into binary file
file:
	rm -rf ./assets/*
	go generate ./assets/...

fmt:
	go fmt ./...

main:
	go build -o bin/main ./cmd/main

test: main gotest

gotest:
	./bin/main
	

ci:
	go test -count=1 -p=1 -v ./tests/...

alltest: gotest ci
	
clean:
	rm -f ./bin/*
