package permutations

import "fmt"

func Permutations(number int) {
	for i := range make([]struct{}, number) {
		fmt.Printf("%d ", i)
	}

	fmt.Println("")
}
