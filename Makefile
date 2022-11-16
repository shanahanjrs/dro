.PHONY: build build-release test

build:
	go build -o dro.local cmd/dro/main.go

build-release:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/dro-Darwin-x86_64 cmd/dro/main.go
	GOOS=linux  GOARCH=amd64 go build -ldflags "-s -w" -o bin/dro-Linux-x86_64 cmd/dro/main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o bin/dro-Darwin-arm64 cmd/dro/main.go
	GOOS=linux  GOARCH=arm64 go build -ldflags "-s -w" -o bin/dro-Linux-arm64 cmd/dro/main.go

test:
	go test ./...

