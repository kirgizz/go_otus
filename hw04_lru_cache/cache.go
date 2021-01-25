package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Capacity  int
	Cache     map[Key]listItem
	Queue     list
	cacheItem cacheItem
}

type cacheItem struct {
	keys []Key
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity, map[Key]listItem{}, list{}, cacheItem{keys: []Key{}}}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	keyExists := false
	var item listItem

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

	l.cacheItem.keys = append(l.cacheItem.keys, key)

	if l.Queue.Len() > l.Capacity {
		delete(l.Cache, l.cacheItem.keys[0])
		l.Queue.Remove(l.Queue.Back())
		l.cacheItem.keys[0] = l.cacheItem.keys[1]
	}
	return keyExists
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if v, ok := l.Cache[key]; ok {
		l.Queue.MoveToFront(&v)
		return v.Value, ok
	}
	return nil, false
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
