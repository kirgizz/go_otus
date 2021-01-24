package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value interface{}
	Next  *listItem
	Prev  *listItem
}

type list struct {
	len   int
	front *listItem
	back  *listItem
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *listItem {
	return l.front
}

func (l *list) Back() *listItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *listItem {
	item := &listItem{
		Value: v,
		Next:  l.front,
	}

	if l.len > 0 {
		l.front.Prev = item
	} else {
		l.back = item
	}

	l.front = item
	l.len++

	return item
}

func (l *list) PushBack(v interface{}) *listItem {
	item := &listItem{
		Value: v,
		Prev:  l.back,
	}

	if l.len > 0 {
		l.back.Next = item
	} else {
		l.front = item
	}

	l.back = item
	l.len++
	return item
}

func (l *list) Remove(i *listItem) {
	prev := i.Prev
	next := i.Next

	if next != nil {
		next.Prev = prev
	} else {
		l.back = prev
	}

	if prev != nil {
		prev.Next = next
	} else {
		l.front = next
	}

	if l.len == 0 {
		return
	}
	l.len--
}

func (l *list) MoveToFront(i *listItem) {
	if i.Prev == nil {
		return
	}

	i.Prev.Next = i.Next

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	i.Prev = nil
	i.Next = l.front

	l.front.Prev = i
	l.front = i
}
