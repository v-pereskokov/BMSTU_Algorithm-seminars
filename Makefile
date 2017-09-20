.PHONY: all test utils permutations list run

all: test

test:
	go test ./...

utils:
	go test ./src/utils/

permutations:
	go test ./src/permutations/

list:
	go test ./src/linkedList/

run:
	go run ./src/main.go
