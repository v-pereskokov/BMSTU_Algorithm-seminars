package linkedList

import (
	"testing"
)

func TestSinglyListInsertAfter(t *testing.T) {
	list := NewSingleList()

	InsertSingleListElement(list, 3)
	InsertSingleListElement(list, 5)
	InsertSingleListElement(list, 6)

	array := []int{3, 5, 6}
	for _, value := range array {
		if value != list.Head.Data {
			t.Errorf("Don't match: %d and %d\n", value, list.Head.Data)
		}

		list.Head = list.Head.Next
	}
}

func TestSinglyListHas(t *testing.T) {
	list := NewSingleList()

	InsertSingleListElement(list, 3)
	InsertSingleListElement(list, 5)
	InsertSingleListElement(list, 6)

	if !HasSingleListElement(list, 3) {
		t.Error("Don't match: not found")
	}
}

func TestSinglyListNotHas(t *testing.T) {
	list := NewSingleList()

	InsertSingleListElement(list, 3)
	InsertSingleListElement(list, 5)
	InsertSingleListElement(list, 6)

	if HasSingleListElement(list, 0) {
		t.Error("Don't match: not found")
	}
}

func TestSinglyListLoop(t *testing.T) {
	list := NewSingleList()

	InsertSingleListElement(list, 3)
	InsertSingleListElement(list, 5)
	InsertSingleListElement(list, 6)

	list.Head.Next.Next.Next = list.Head

	if !IsLoopListSingleList(list) {
		t.Error("Don't match: loop not found")
	}
}

func TestSinglyListNoLoop(t *testing.T) {
	list := NewSingleList()

	InsertSingleListElement(list, 3)
	InsertSingleListElement(list, 5)
	InsertSingleListElement(list, 6)

	if IsLoopListSingleList(list) {
		t.Error("Don't match: loop found")
	}
}

func TestSinglyListRemove(t *testing.T) {
	list := NewSingleList()

	InsertSingleListElement(list, 3)
	InsertSingleListElement(list, 5)
	InsertSingleListElement(list, 6)

	RemoveSingleListElement(list, 5)

	array := []int{3, 6}
	for _, value := range array {
		if value != list.Head.Data {
			t.Logf("%v\n", list)
			t.Errorf("Don't match: %d and %d\n", value, list.Head.Data)
		}

		list.Head = list.Head.Next
	}
}

func TestSinglyListNotRemove(t *testing.T) {
	list := NewSingleList()

	InsertSingleListElement(list, 3)
	InsertSingleListElement(list, 5)
	InsertSingleListElement(list, 6)

	RemoveSingleListElement(list, 7)

	array := []int{3, 5, 6}
	for _, value := range array {
		if value != list.Head.Data {
			t.Logf("%v\n", list)
			t.Errorf("Don't match: %d and %d\n", value, list.Head.Data)
		}

		list.Head = list.Head.Next
	}
}

func TestSinglyListRemoveEmpty(t *testing.T) {
	list := NewSingleList()

	RemoveSingleListElement(list, 7)

	if list.Len > 0 {
		t.Errorf("Don't match: size more than 0\n")
	}
}
