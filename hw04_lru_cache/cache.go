package hw04_lru_cache //nolint:golint,stylecheck
import "fmt"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Capacity int
	Cache    map[Key]listItem
	Queue    list
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity, map[Key]listItem{}, list{}}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	keyExists := false
	var item listItem

	if l.Queue.Len() >= l.Capacity {
		itemToRemove := l.Queue.Back()
		key := l.FindItemByValue(*itemToRemove)

		fmt.Println("delete key: ", itemToRemove.Value)
		fmt.Println("Before delete ", l.Cache)
		delete(l.Cache, key)
		fmt.Println("After delete ", l.Cache)
		l.Queue.Remove(l.Queue.Back())
	}

	// update key and push to front
	if v, ok := l.Cache[key]; ok {
		keyExists = true
		item = v
		item.Value = value
		l.Queue.MoveToFront(&item)
		l.Cache[key] = item
	} else {
		item = *l.Queue.PushFront(value)
		l.Cache[key] = item
	}

	/*
	fmt.Printf("Queue len %d\n", l.Queue.Len())
	//fmt.Println("Front elem ", l.Queue.Front())
	//fmt.Println("Back elem ", l.Queue.Back())
	fmt.Println("#########")
	for i := l.Queue.Front(); i != nil; i = i.Next {
		fmt.Println(i.Value)
	}

	 */
	return keyExists
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	keyExists := false
	var item listItem
	if v, ok := l.Cache[key]; ok {
		keyExists = true
		item = v
		l.Queue.MoveToFront(&item)
	}
	return item.Value, keyExists
}

func (l *lruCache) Clear() {
	l.Cache = map[Key]listItem{}
	l.Queue = list{}
}

func (l *lruCache) FindItemByValue(item listItem) Key {
	for k, v := range l.Cache {
		if v == item {
			return k
		}
	}
	return ""
}
