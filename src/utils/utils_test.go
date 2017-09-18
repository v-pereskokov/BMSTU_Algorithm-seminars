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
