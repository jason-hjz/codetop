package _46_LRUCache

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	LRU := initLRUCache(4)
	LRU.put(1, 1)
	LRU.put(2, 2)
	LRU.put(3, 3)
	LRU.put(4, 4)
	LRU.put(5, 5)
	fmt.Println(LRU.get(1))
	fmt.Println(LRU.get(5))
	return
}
