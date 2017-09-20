package linkedList

type SinglyListItem struct {
	Data interface{}
	Next *SinglyListItem
}

type SinglyList struct {
	Len  int
	Head *SinglyListItem
	Tail *SinglyListItem
}

func NewSingleList() *SinglyList {
	list := &SinglyList{}

	return list
}

func InsertSingleListElement(list *SinglyList, value interface{}) {
	item := &SinglyListItem{value, nil}

	if list.Head == nil {
		list.Head = item
		list.Tail = item
	} else {
		list.Tail.Next = item
		list.Tail = item
	}

	list.Len++
}

func RemoveSingleListElement(list *SinglyList, value interface{}) {
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

func HasSingleListElement(list *SinglyList, value interface{}) bool {
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

func IsLoopListSingleList(list *SinglyList) bool {
	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
