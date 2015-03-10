cwd=$(shell pwd)

fmt:
	GOPATH=${cwd}:${GOPATH} gofmt -w cmd.go
	GOPATH=${cwd}:${GOPATH} gofmt -w cmd_test.go
