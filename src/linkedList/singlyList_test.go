package linkedList

import "testing"

func TestSinglyListInsertAfter(t *testing.T) {
	list := New()
	InsertAfter(list, 3)
	InsertAfter(list, 5)

	array := []int{0, 5, 3}
	for _, value := range array {
		if value != list.Head.Data {
			t.Errorf("Don't match: %d and %d\n", value, list.Head.Data)
		}

		list.Head = list.Head.Next
	}
}
