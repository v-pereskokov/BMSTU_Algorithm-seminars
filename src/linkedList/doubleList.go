package linkedList

import "fmt"

type DoubleList struct {
	Head *DoubleListItem
	Tail *DoubleListItem
	Len  int
}

type DoubleListItem struct {
	Data interface{}
	Next *DoubleListItem
	Prev *DoubleListItem
}

func NewDoubleList() *DoubleList {
	list := &DoubleList{}
	list.Len = 0
	return list
}

func InsertDoubleListElement(list *DoubleList, value interface{}) {
	newItem := &DoubleListItem{value, list.Head, list.Tail}

	if list.Head == nil {
		list.Head = newItem
		list.Tail = newItem
	} else {
		list.Head = newItem
		list.Head.Prev = newItem
		list.Tail.Next = newItem
	}

	list.Len++
}

func HasDoubleListElement(list *DoubleList, value interface{}) bool {
	current := list.Head

	for {
		if current == nil {
			return false
		}

		if current.Data == value {
			return true
		}

		current = current.Next

		if current == list.Tail && current.Data != value {
			return false
		}
	}
}
