package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	head  *ListItem
	tail  *ListItem
	items map[*ListItem]bool
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	li := ListItem{Value: v, Prev: nil, Next: l.head}
	if l.head != nil {
		l.head.Prev = &li
	}
	if l.tail == nil {
		// if it is first push
		l.tail = &li
	}
	l.head = &li
	l.len++
	l.items[l.head] = true
	return &li
}

func (l *list) PushBack(v interface{}) *ListItem {
	li := ListItem{Value: v, Prev: l.tail, Next: nil}

	if l.tail != nil {
		l.tail.Next = &li
	}
	if l.head == nil {
		// if it is first push
		l.head = &li
	}
	l.tail = &li
	l.len++
	l.items[l.tail] = true
	return &li
}

func (l *list) Remove(li *ListItem) {
	// вообще нужно проверять что элемент, который удаляем вообще находится в этом списке
	// а не только проверять не больше ли нуля список (там утечки памяти, не очевидное поведение и тд и тп)
	// я не знаю что быстрее использовать ссылку на объект в качестве ключа мапы или таки обхот всего списка
	// решил попровбовать пока так
	val, ok := l.items[li]
	if ok && val {
		if l.head == li {
			l.head = li.Next
		}
		if l.tail == li {
			l.tail = li.Prev
		}
		if li.Next != nil {
			li.Next.Prev = li.Prev
		}
		if li.Prev != nil {
			li.Prev.Next = li.Next
		}
		l.len--
		delete(l.items, li)
	}
}

func (l *list) MoveToFront(li *ListItem) {
	l.Remove(li)
	l.PushFront(li.Value)
}

func NewList() List {
	return &list{0, nil, nil, make(map[*ListItem]bool)}
}
