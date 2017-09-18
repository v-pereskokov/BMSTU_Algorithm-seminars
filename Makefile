.PHONY: all test range run

all: test

test:
	go test ./...

range:
	go test ./src/utils/

run:
	go run ./src/main.go
