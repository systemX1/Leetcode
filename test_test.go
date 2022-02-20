package leetcode

import (
	"container/heap"
	"fmt"
	"log"
	"sort"
	"testing"
)

func init() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
}

func TestMap(t *testing.T) {
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	log.Println(pairs['('])
	log.Println(pairs[')'])

}

type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}
type StuScores []StuScore

func (s StuScores) Len() int {
	return len(s)
}
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestSort(t *testing.T) {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}
	fmt.Println("Default:\n\t", stus)
	sort.Sort(stus)
	fmt.Println("Is Sorted?\n\t", sort.IsSorted(stus))
	fmt.Println("Sorted:\n\t", stus)
	sort.Sort(sort.Reverse(stus))
	fmt.Println("Reverse Sorted:\n\t", stus)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	fmt.Printf("heap: %v\n", *h)
	heap.Push(h, 3)
	heap.Pop(h)
	fmt.Printf("heap: %v\n", *h)
}

func TestValidateStackSequences(t *testing.T) {
	fn := func(pushed []int, popped []int) bool {
		l := len(pushed)
		stack := []int{}
		i, j := 0, 0
		for i < l {
			stack = append(stack, pushed[i])
			for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
				log.Println(stack, i, j, "LINE 30")
				stack = stack[:len(stack)-1]
				j++
			}
			log.Println(stack, i, j, "LINE 34")
			i++
		}
		return len(stack) == 0
	}
	fn([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
}

func TestCal(t *testing.T) {
	log.Println(-5/4, "|", -4/4, -3/4, -2/4, -1/4, "|", 0/4, 1/4)
	var a, b int
	a, b = 2147483647, -2147483640
	log.Println(a - -1, b-100)
}

func TestValidSudoku(t *testing.T) {
	getHtIdx := func(i, j int) int {
		return (i/3)*3 + j/3
	}
	fn := func(board [][]byte) bool {
		row, col, ht := [9][9]int{}, [9][9]int{}, [9][9]int{}
		for i, r := range board {
			for j, ch := range r {
				if ch == '.' {
					continue
				}
				elem := ch - '1'
				idx := getHtIdx(i, j)
				if row[i][elem] != 0 || col[j][elem] != 0 || ht[idx][elem] != 0 {
					return false
				}
				row[i][elem], col[j][elem], ht[idx][elem] = 1, 1, 1
			}
		}
		return true
	}
	board := [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'5', '.', '.', '.', '.', '.', '.', '9', '.'}, {'.', '.', '.', '5', '6', '.', '.', '.', '.'}, {'4', '.', '3', '.', '.', '.', '.', '.', '1'}, {'.', '.', '.', '7', '.', '.', '.', '.', '.'}, {'.', '.', '.', '5', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	log.Println(fn(board))
}

// type Node struct {
// 	val  int
// 	key  int
// 	pre  *Node
// 	next *Node
// }

// type LRUCache struct {
// 	ht   map[int]*Node
// 	head *Node
// 	tail *Node
// 	cap  int
// 	size int
// }

// func LRUCacheConstructor(capacity int) LRUCache {
// 	lru := LRUCache{cap: capacity, size: 0}
// 	lru.ht = make(map[int]*Node)
// 	lru.head = &Node{}
// 	lru.tail = &Node{}
// 	lru.head.next = lru.tail
// 	lru.tail.pre = lru.head
// 	return lru
// }

// func (this *LRUCache) Get(key int) int {
// 	if node, ok := this.ht[key]; ok {
// 		this.Refresh(node)
// 		return node.val
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	var node *Node
// 	var ok bool
// 	if node, ok = this.ht[key]; ok {
// 		node.val = value
// 	} else {
// 		if this.size == this.cap {
// 			del := this.head.next
// 			del.next.pre = del.pre
// 			del.pre.next = del.next
// 			delete(this.ht, del.key)
// 			this.size--
// 		}
// 		node = &Node{val: value, key: key}
// 		this.ht[key] = node
// 		this.size++
// 	}
// 	this.Refresh(node)
// }

// func (this *LRUCache) Refresh(node *Node) {
// 	if node.next != nil {
// 		node.next.pre = node.pre
// 		node.pre.next = node.next
// 	}
// 	node.next = this.tail
// 	node.pre = this.tail.pre
// 	node.pre.next = node
// 	this.tail.pre = node
// }

// func (this *LRUCache) Debug() string {
// 	var buf bytes.Buffer
// 	buf.WriteByte('{')
// 	defer buf.WriteByte('}')

// 	buf.WriteString(fmt.Sprintf("%v %v %v", this.cap, this.size, this.ht))
// 	for p := this.head; p != nil; p = p.next {
// 		buf.WriteString(fmt.Sprintf(" %p:%v ", p, p.val))
// 	}
// 	return buf.String()
// }

// func TestLRUCache(t *testing.T) {
// 	lru := LRUCacheConstructor(2)
// 	log.Println(lru.Debug())
// 	lru.Put(1, 1)
// 	log.Println(lru.Debug())
// 	lru.Put(2, 2)
// 	log.Println(lru.Debug())
// 	lru.Get(1)
// 	log.Println(lru.Debug())
// 	lru.Put(3, 3)
// 	log.Println(lru.Debug())
// 	lru.Get(2)
// 	log.Println(lru.Debug())

// }

// func TestDesignCircularQueue(t *testing.T) {
// 	// cq := MyCircularQueueConstructor(3)
// 	// log.Println(cq.EnQueue(1))
// 	// log.Println(cq)
// 	// log.Println(cq.EnQueue(2))
// 	// log.Println(cq)
// 	// log.Println(cq.EnQueue(3))
// 	// log.Println(cq)
// 	// log.Println(cq.EnQueue(4))
// 	// log.Println(cq)
// 	// log.Println(cq.Rear())
// 	// log.Println(cq.Front())
// 	// log.Println(cq.IsEmpty())
// 	// log.Println(cq.IsFull())
// 	// log.Println(cq.DeQueue())
// 	// log.Println(cq)
// 	// log.Println(cq.EnQueue(4))
// 	// log.Println(cq)
// 	// log.Println(cq.Rear())
// 	// log.Println(cq.Front())
// 	// log.Println(cq.IsEmpty())
// 	// log.Println(cq.IsFull())
// 	// log.Println(cq.DeQueue())

// 	cq := MyCircularQueueConstructor(6)
// 	log.Println(cq.EnQueue(6))
// 	log.Println(cq)
// 	log.Println(cq.Rear())
// 	log.Println(cq.Front())
// 	log.Println(cq.DeQueue())
// 	log.Println(cq)
// 	log.Println(cq.EnQueue(5), "EnQueue(5)")
// 	log.Println(cq)
// 	log.Println(cq.Rear())
// 	log.Println(cq.DeQueue(), "DeQueue()")
// 	log.Println(cq)
// 	log.Println(cq.Front())
// 	log.Println(cq.DeQueue())
// 	log.Println(cq)
// 	log.Println(cq.DeQueue())
// 	log.Println(cq)
// 	log.Println(cq.DeQueue())
// 	log.Println(cq)
// }

// func TestImplementQueueUsingStacks(t *testing.T) {
// 	q := MyQueueConstructor()
// 	q.Push(1)
// 	log.Println(q)
// 	q.Push(2)
// 	log.Println(q)
// 	q.Pop()
// 	log.Println(q)
// 	log.Println("len(this.s1)", len(q.s1), len(q.s1) != 0)
// 	q.Push(3)
// 	log.Println("len(this.s1)", len(q.s1))
// 	log.Println(q)
// 	q.Push(4)
// 	log.Println(q)
// 	q.Pop()
// 	log.Println(q)
// 	log.Println(q.Peek())
// }

// func TestImplementStackUsingQueues(t *testing.T) {
// 	s := MyStackConstructor()
// 	s.Push(1)
// 	log.Println(s)
// 	s.Push(2)
// 	log.Println(s)
// 	log.Println(s.Top())
// 	log.Println(s.Pop())
// 	log.Println(s)
// 	log.Println(s.Empty())

// }
