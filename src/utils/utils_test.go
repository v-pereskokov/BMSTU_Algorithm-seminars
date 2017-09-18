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
	list, err := Reverse([]int{4, 2, 4, 0, 0, 1}, -1, -1)
	if err != nil {
		t.Errorf("Error reverse: %v", err)
	}

	length := len(list)
	testArray := []int{4, 2, 4, 0, 0, 1}

	for i, value := range list {
		if testArray[length-i-1] != value {
			t.Errorf("Don't match: %d and %d\n", list[length-i-1], value)
		}
	}
}

func TestReverseEmpty(t *testing.T) {
	list, err := Reverse([]int{}, -1, -1)
	if err != nil {
		t.Errorf("Error reverse: %v", err)
	}

	testArray := []int{}

	if len(testArray) != len(list) {
		t.Errorf("Don't empty: %v and %v\n", list, testArray)
	}
}

func TestReverseRange(t *testing.T) {
	begin := 2
	end := 4

	list, err := Reverse([]int{4, 2, 4, 0, 0, 1}, begin, end)
	if err != nil {
		t.Errorf("Error reverse: %v", err)
	}

	length := len(list)
	testArray := []int{4, 2, 4, 0, 0, 1}

	for i := begin; i < end; i++ {
		t.Logf("%d and %d\n", testArray[length-i], list[i])
		if testArray[length-i] != list[i] {
			t.Errorf("Don't match: %d and %d\n", testArray[length-i], list[i])
		}
	}
}

func TestReverseRangeError(t *testing.T) {
	begin := 4
	end := 2

	_, err := Reverse([]int{4, 2, 4, 0, 0, 1}, begin, end)
	if err == nil {
		t.Error("Error reverse: begin > end")
	}
}
