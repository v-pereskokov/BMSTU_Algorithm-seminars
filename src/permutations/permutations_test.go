package permutations

import "testing"

func TestPermutations(t *testing.T) {
	checkArray := []int{
		1, 2, 3,
		1, 3, 2,
		2, 1, 3,
		2, 3, 1,
		3, 1, 2,
		3, 2, 1,
	}
	checkIndex := 0

	for permutation := range Permutations(3) {
		t.Log(permutation)
		for _, value := range permutation {
			if value != checkArray[checkIndex] {
				t.Errorf("Don't match: %d and %d \n", value, checkArray[checkIndex])
			}

			checkIndex++
		}
	}
}
