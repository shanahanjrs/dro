.PHONY: build build-release test

build:
	go build -o dro

build-release:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/dro-macos-amd64 cmd/dro/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/dro-linux-amd64 cmd/dro/main.go

test:
	go test ./...

