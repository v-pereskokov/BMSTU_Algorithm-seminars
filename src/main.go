package main

import (
	"fmt"
	"github.com/vladpereskokov/BMSTU_Algorithm-seminars/src/permutations"
)

func main() {
	for v := range permutations.Permutations(3) {
		fmt.Println(string(v))
	}
}
