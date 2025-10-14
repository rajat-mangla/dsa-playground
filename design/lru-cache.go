package design

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

type LRUCache struct {
	capacity  int
	hashStore map[int]*Node
	head      *Node
	tail      *Node
}

func Constructor(capacity int) LRUCache {
	tail := &Node{}
	head := &Node{}
	tail.prev = head
	head.next = tail

	return LRUCache{
		capacity:  capacity,
		hashStore: map[int]*Node{},
		head:      head,
		tail:      tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hashStore[key]
	if ok {
		deleteNode(node)
		addToHead(this.head, node)

		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hashStore[key]
	if ok {
		node.value = value
		deleteNode(node)
		addToHead(this.head, node)
		return
	}

	if len(this.hashStore) == this.capacity {
		tailPrev := this.tail.prev
		delete(this.hashStore, tailPrev.key)
		deleteNode(tailPrev)
	}

	newNode := &Node{
		key:   key,
		value: value,
	}
	this.hashStore[key] = newNode
	addToHead(this.head, newNode)
}

func deleteNode(node *Node) {
	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
	node.next = nil
	node.prev = nil
}

func addToHead(head, node *Node) {
	firstNode := head.next
	head.next = node
	node.prev = head
	node.next = firstNode
	firstNode.prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
