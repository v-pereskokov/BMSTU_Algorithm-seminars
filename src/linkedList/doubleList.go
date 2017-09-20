package linkedList

type List struct {
	Head *Item
	Tail *Item
	Len  int
}

type Item struct {
	Data interface{}
	Next *Item
	Prev *Item
}

func NewDoubleList() *List {
	list := &List{}
	list.Len = 0
	return list
}

func InsertDoubleListElement(list *List, value interface{}) *List {
	newItem := &Item{value, list.Head, list.Tail}

	if list.Head == nil {
		list.Head = newItem
		list.Tail = newItem
	} else {
		list.Head = newItem
		list.Head.Prev = newItem
		list.Tail.Next = newItem
	}

	list.Len++

	return list
}

func HasDoubleListElement(list *List, value interface{}) bool {
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

func RemoveDoubleListElement(list *List, value interface{}) *List {
	if list.Head == nil {
		return list
	}

	first := list.Head
	Tail := list.Tail

	for {
		if Tail.Next == nil {
			return list
		}

		if first.Data == value {
			first.Prev.Next = first.Next
			first.Next.Prev = first.Prev

			first.Prev = nil
			first.Next = nil
			first.Data = nil

			list.Len--
			return list
		} else {
			first = first.Next
		}
	}
}
