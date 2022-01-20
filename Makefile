.PHONY: build clean fmt test

default: build

clean:
	@go mod tidy && rm -rf nebula-diag-cli

fmt:
	@find . -type f -iname \*.go -exec go fmt {} \;

build: clean fmt
	@go build -o nebula-diag-cli cmd/main.go

test: build
	go test -v ./tests/...