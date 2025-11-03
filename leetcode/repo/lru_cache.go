package repo

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*doubleList
	head     *doubleList
	tail     *doubleList
}
type doubleList struct {
	prev *doubleList
	next *doubleList
	key  int
	val  int
}

func Constructor(capacity int) LRUCache {
	h := &doubleList{}
	t := &doubleList{}
	h.next = t
	t.prev = h
	return LRUCache{
		capacity: capacity,
		head:     h,
		tail:     t,
		cache:    make(map[int]*doubleList),
	}
}

func (this *LRUCache) moveToHead(node *doubleList) {
	this.removeNode(node)
	this.insertToHead(node)
}

func (this *LRUCache) removeNode(node *doubleList) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) insertToHead(node *doubleList) {
	this.head.next.prev = node
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.val = value
		this.moveToHead(node)
		return
	}

	if this.size == this.capacity {
		todelete := this.tail.prev
		delete(this.cache, todelete.key)
		this.removeNode(todelete)
		this.size--
	}

	newNode := &doubleList{
		key: key,
		val: value,
	}
	this.insertToHead(newNode)
	this.cache[key] = newNode
	this.size++
}
