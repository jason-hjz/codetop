package _46_LRUCache

//func main() {
//	LRU := initLRUCache(4)
//	LRU.put(1, 1)
//	LRU.put(2, 2)
//	LRU.put(3, 3)
//	LRU.put(4, 4)
//	LRU.put(5, 5)
//	fmt.Println(LRU.get(1))
//	fmt.Println(LRU.get(5))
//	return
//}

type DLinkedNode struct {
	key, value int
	next, prev *DLinkedNode
}

type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

// 摘除节点、加入头结点、置顶节点、删除尾节点、初始化新节点、初始化cache、get、put

func (cache *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (cache *LRUCache) addToHead(node *DLinkedNode) {
	node.next = cache.head.next
	cache.head.next.prev = node
	cache.head.next = node
	node.prev = cache.head
}

func (cache *LRUCache) topNode(node *DLinkedNode) {
	cache.removeNode(node)
	cache.addToHead(node)
}

func initNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func initLRUCache(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		size:     0,
		cache:    map[int]*DLinkedNode{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (cache *LRUCache) removeTail() *DLinkedNode {
	node := cache.tail.prev
	cache.removeNode(node)
	return node
}

func (cache *LRUCache) get(key int) int {
	if _, ok := cache.cache[key]; !ok {
		return -1
	}
	node := cache.cache[key]
	cache.addToHead(node)
	return node.value
}

func (cache *LRUCache) put(key int, value int) {
	if _, ok := cache.cache[key]; !ok {
		node := initNode(key, value)
		cache.cache[key] = node
		cache.addToHead(initNode(key, value))
		cache.size++
		if cache.size > cache.capacity {
			tail := cache.removeTail()
			delete(cache.cache, tail.key)
			cache.size--
		}
	}
	cache.cache[key].value = value
	cache.addToHead(cache.cache[key])
}
