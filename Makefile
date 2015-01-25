cwd=$(shell pwd)

all:
	GOPATH=${cwd}:${GOPATH} go build -o bin/example example.go

test:
	GOPATH=${cwd}:${GOPATH} go test cmd

fmt:
	GOPATH=${cwd}:${GOPATH} gofmt -w example.go
	GOPATH=${cwd}:${GOPATH} gofmt -w src/cmd/cmd.go
	GOPATH=${cwd}:${GOPATH} gofmt -w src/cmd/cmd_test.go

clean:
	rm bin/*
