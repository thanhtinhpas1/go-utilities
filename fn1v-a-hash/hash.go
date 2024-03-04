package main

import "bytes"

type Counter struct {
	items []CounterItem
	size  int // number of active items in item slice
}

type CounterItem struct {
	Key   []byte // key, nil in Counter.items if not used
	Count int    // number of occurences
}

// This part is core of FN1v hash
const (
	// FNV-1 64-bit constants from hash/fnv
	offset64 = 14695981039346656037
	prime64  = 1099511628211

	// initial length for counter
	initialLength = 1024
)

func (c *Counter) Incr(key []byte, n int) {
	hash := uint64(offset64)
	for _, c := range key {
		hash *= prime64
		hash ^= uint64(c)
	}

	// make 64-bit hash in range for items slice
	index := int(hash & uint64(len(c.items)-1))

	// if current items more than half full, double length and reinsert items
	if c.size >= len(c.items)/2 {
		newLen := len(c.items) * 2
		if newLen == 0 {
			newLen = initialLength
		}
		newC := Counter{items: make([]CounterItem, newLen), size: newLen}
		for _, item := range c.items {
			if item.Key != nil {
				newC.Incr(item.Key, item.Count)
			}
		}
		c.items = newC.items
		index = int(hash & uint64(len(newC.items)-1))
	}

	for {
		if c.items[index].Key == nil {
			// found empty slot, add new item (copying key)
			keyCopy := make([]byte, len(key))
			copy(keyCopy, key)
			c.items[index] = CounterItem{keyCopy, n}
			c.size++
			return
		}

		if bytes.Equal(c.items[index].Key, key) {
			// found matching slot, increment existing count.
			c.items[index].Count += n
			return
		}

		index++
		if index >= len(c.items) {
			index = 0
		}
	}
}

// Items returns a copy of the incremented items.
func (c *Counter) Items() []CounterItem {
	var items []CounterItem
	for _, item := range c.items {
		if item.Key != nil {
			items = append(items, item)
		}
	}
	return items
}
