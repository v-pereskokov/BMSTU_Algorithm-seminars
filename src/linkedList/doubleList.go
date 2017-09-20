package linkedList

type DoubleListItem struct {
	Data interface{}
	Next *DoubleListItem
}

type DoubleList struct {
	Len  int
	Head *DoubleListItem
	Tail *DoubleListItem
}

func NewDoubleList() *DoubleList {
	list := &DoubleList{}

	return list
}

func InsertDoubleListElement(list *DoubleList, value interface{}) {
	item := &DoubleListItem{value, nil}

	if list.Head == nil {
		list.Head = item
		list.Tail = item
	} else {
		list.Tail.Next = item
		list.Tail = item
	}

	list.Len++
}

func RemoveDoubleListElement(list *DoubleList, value interface{}) {
	if list.Head == nil {
		return
	}

	current := list.Head
	prev := list.Head
	for {
		if current == nil {
			return
		}

		if current.Data == value {
			prev.Next = current.Next
			current.Next = nil

			list.Len--

			return
		} else {
			prev = current
			current = current.Next
		}
	}
}

func HasDoubleListElement(list *DoubleList, value interface{}) bool {
	for {
		if list.Head == nil {
			return false
		}

		if list.Head.Data == value {
			return true
		}

		list.Head = list.Head.Next
	}
}
