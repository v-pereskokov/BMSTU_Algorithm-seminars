package utils

import "testing"

func TestRangeOk(t *testing.T) {
	list := Range(4)

	for i := 0; i < 4; i++ {
		if list[i] != i+1 {
			t.Errorf("Don't match: %d and %d\n", list[i], i+1)
		}
	}
}

func TestRangeFail(t *testing.T) {
	list := Range(-1)

	if len(list) > 0 {
		t.Errorf("Error init: %d\n", len(list))
	}
}

func TestReverse(t *testing.T) {
	list := Reverse([]int{4, 2, 4, 0, 0, 1})
	length := len(list)

	testArray := []int{4, 2, 4, 0, 0, 1}

	for i, value := range list {
		if testArray[length-i-1] != value {
			t.Errorf("Don't match: %d and %d\n", list[length-i-1], value)
		}
	}
}

func TestReverseEmpty(t *testing.T) {
	list := Reverse([]int{})
	testArray := []int{}

	if len(testArray) != len(list) {
		t.Errorf("Don't empty: %v and %v\n", list, testArray)
	}
}
