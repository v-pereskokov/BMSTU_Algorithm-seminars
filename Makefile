.PHONY: all test stack run

all: test

test:
	go test ./...

stack:
	go test ./src/stack/

run:
	go run ./src/main.go
