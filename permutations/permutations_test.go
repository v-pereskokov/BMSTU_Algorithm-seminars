package permutations

import "testing"

func TestPermutations(t *testing.T) {
	checkArray := []string{
		"[1 2 3]",
		"[1 3 2]",
		"[2 1 3]",
		"[2 3 1]",
		"[3 1 2]",
		"[3 2 1]",
	}
	checkIndex := 0

	for permutation := range Permutations(3) {
		t.Log(permutation)
		if string(permutation) != checkArray[checkIndex] {
			t.Errorf("Don't match: %s and %s \n", string(permutation), checkArray[checkIndex])
		}

		checkIndex++
	}
}
