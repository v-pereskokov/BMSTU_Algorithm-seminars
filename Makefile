.PHONY: all test utils permutations

all: test

test: utils permutations

utils:
	go test ./utils/

permutations:
	go test ./permutations/
