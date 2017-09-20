package linkedList

import "testing"

func TestDoubleListInsert(t *testing.T) {
	list := NewDoubleList()

	InsertDoubleListElement(list, 3)
	InsertDoubleListElement(list, 5)
	InsertDoubleListElement(list, 6)

	t.Logf("%v\n", list.Head)
	array := []int{6, 5, 3}
	for _, value := range array {
		if value != list.Head.Data {
			t.Errorf("Don't match: %d and %d\n", value, list.Head.Data)
		}

		list.Head = list.Head.Next
	}
}
