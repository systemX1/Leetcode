/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */
package leetcode

// @lc code=start
type Node struct {
	val  int
	key  int
	pre  *Node
	next *Node
}

type LRUCache struct {
	ht   map[int]*Node
	head *Node
	tail *Node
	cap  int
	size int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{cap: capacity, size: 0}
	lru.ht = make(map[int]*Node)
	lru.head = &Node{}
	lru.tail = &Node{}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.ht[key]; ok {
		this.Refresh(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	var node *Node
	var ok bool
	if node, ok = this.ht[key]; ok {
		node.val = value
	} else {
		if this.size == this.cap {
			del := this.head.next
			del.next.pre = del.pre
			del.pre.next = del.next
			delete(this.ht, del.key)
			this.size--
		}
		node = &Node{val: value, key: key}
		this.ht[key] = node
		this.size++
	}
	this.Refresh(node)
}

func (this *LRUCache) Refresh(node *Node) {
	if node.next != nil {
		node.next.pre = node.pre
		node.pre.next = node.next
	}
	node.next = this.tail
	node.pre = this.tail.pre
	node.pre.next = node
	this.tail.pre = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
