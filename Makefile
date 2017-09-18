.PHONY: all test utils run

all: test

test:
	go test ./...

utils:
	go test ./src/utils/

permutations:
	go test ./src/permutations/

run:
	go run ./src/main.go
