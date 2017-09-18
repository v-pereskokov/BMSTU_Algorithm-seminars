package linkedList

import (
	"testing"
)

func TestSinglyListInsertAfter(t *testing.T) {
	list := New()

	Insert(list, 3)
	Insert(list, 5)
	Insert(list, 6)

	array := []int{3, 5, 6}
	for _, value := range array {
		if value != list.Head.Data {
			t.Errorf("Don't match: %d and %d\n", value, list.Head.Data)
		}

		list.Head = list.Head.Next
	}
}
