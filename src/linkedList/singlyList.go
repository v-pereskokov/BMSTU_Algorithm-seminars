package linkedList

type SinglyListItem struct {
	Data interface{}
	Next *SinglyListItem
}

type SinglyList struct {
	Len  int
	Head *SinglyListItem
}

func New() *SinglyList {
	list := &SinglyList{0, &SinglyListItem{0, nil}}

	return list
}

func InsertAfter(list *SinglyList, value interface{}) {
	item := &SinglyListItem{value, list.Head.Next}

	item.Next = list.Head.Next
	list.Head.Next = item

	list.Len++
}
