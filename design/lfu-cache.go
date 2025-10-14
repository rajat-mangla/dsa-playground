package design

type LFUNode struct {
	key  int
	val  int
	freq int
	prev *LFUNode
	next *LFUNode
}

type CacheLL struct {
	size int
	head *LFUNode
	tail *LFUNode
}

type LFUCache struct {
	keyStore  map[int]*LFUNode
	freqStore map[int]*CacheLL
	minFreq   int
	cap       int
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		cap:       capacity,
		keyStore:  map[int]*LFUNode{},
		freqStore: map[int]*CacheLL{},
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.keyStore[key]
	if !ok {
		return -1
	}

	this.moveNode(node, node.val)
	return node.val
}

func (this *LFUCache) Put(key int, value int) {
	node, ok := this.keyStore[key]
	if ok {
		this.moveNode(node, value)
		return
	}

	if len(this.keyStore) == this.cap {
		lruCache, _ := this.freqStore[this.minFreq]

		node = lruCache.tail.prev
		removeNode(node)
		lruCache.size -= 1
		if lruCache.size == 0 {
			delete(this.freqStore, this.minFreq)
		}
		delete(this.keyStore, node.key)
	}

	node = &LFUNode{
		key:  key,
		val:  value,
		freq: 1,
	}

	// add node to freq 1
	lruCache, ok := this.freqStore[node.freq]
	if !ok {
		head := &LFUNode{}
		tail := &LFUNode{}
		head.next = tail
		tail.prev = head
		lruCache = &CacheLL{
			head: head,
			tail: tail,
		}
	}

	addNodeToHead(lruCache.head, node)
	this.keyStore[key] = node

	lruCache.size += 1
	this.freqStore[node.freq] = lruCache
	this.minFreq = 1
}

func (this *LFUCache) moveNode(node *LFUNode, value int) {
	freq := node.freq
	// remove node from old freq
	lruCache, ok := this.freqStore[freq]
	lruCache.size -= 1
	removeNode(node)
	if lruCache.size == 0 {
		delete(this.freqStore, freq)
		if this.minFreq == freq {
			this.minFreq = freq + 1
		}
	}

	// add node to new freq
	freq += 1
	node.freq = freq
	node.val = value
	// add node to new freq
	lruCache, ok = this.freqStore[freq]
	if !ok {
		head := &LFUNode{}
		tail := &LFUNode{}
		head.next = tail
		tail.prev = head
		lruCache = &CacheLL{
			head: head,
			tail: tail,
		}
	}
	lruCache.size += 1
	addNodeToHead(lruCache.head, node)

	this.freqStore[freq] = lruCache
}

func removeNode(node *LFUNode) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
	node.next = nil
	node.prev = nil
}

func addNodeToHead(head, node *LFUNode) {
	nextNode := head.next

	node.prev = head
	node.next = nextNode

	head.next = node
	nextNode.prev = node
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
