package linkedList

type SinglyListItem struct {
	data interface{}
	next *SinglyListItem
	list *SinglyList
}

func (item SinglyListItem) Data() interface{} {
	return item.data
}

func (item SinglyListItem) Next() *SinglyListItem {
	return item.next
}

func (item SinglyListItem) List() *SinglyList {
	return item.list
}

type SinglyList struct {
	len  int
	head *SinglyListItem
}

func (list SinglyList) Len() int {
	return list.len
}

func (list SinglyList) Head() *SinglyListItem {
	return list.head
}

func New() *SinglyList {
	list := new(SinglyList)
	list.len = 0

	return list
}
