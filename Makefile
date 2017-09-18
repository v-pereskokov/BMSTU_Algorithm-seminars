.PHONY: all test utils permutations linkedList run

all: test

test:
	go test ./...

utils:
	go test ./src/utils/

permutations:
	go test ./src/permutations/

linkedList:
	go test ./src/linkedList/

run:
	go run ./src/main.go
