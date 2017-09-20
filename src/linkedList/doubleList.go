package linkedList

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
	if list.Head == nil {
		return false
	}
	first := list.Head

	for {
		if first.Data == value {
			return true
		} else {
			if first.Next != nil {
				first = first.Next
			} else {
				return false
			}
		}
	}

	return false
}

func RemoveDoubleListElement(list *DoubleList, value interface{}) {
	if list.Head == nil {
		return
	}

	first := list.Head
	Tail := list.Tail

	for {
		if Tail.Next == nil {
			return
		}

		if first.Data == value {
			first.Prev.Next = first.Next
			first.Next.Prev = first.Prev

			first.Prev = nil
			first.Next = nil
			first.Data = nil

			list.Len--

			return
		} else {
			first = first.Next
		}
	}
}
