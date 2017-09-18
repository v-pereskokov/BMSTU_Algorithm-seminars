.PHONY: all test utils run

all: test

test:
	go test ./...

utils:
	go test ./src/utils/

run:
	go run ./src/main.go
