.PHONY: all test stackTest run

all: test

test: stackTest

stackTest:
	go test ./src/stack/

run:
	go run ./src/main.go
