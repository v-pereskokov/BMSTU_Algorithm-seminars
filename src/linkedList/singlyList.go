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

func New() *SinglyList {
	list := &SinglyList{}

	return list
}

func Insert(list *SinglyList, value interface{}) {
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

func Has(list *SinglyList, value interface{}) bool {
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
