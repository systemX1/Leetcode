/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */
package leetcode

// @lc code=start
type MyCircularQueue struct {
	s       []int
	front   int
	length  int
	capcity int
}

func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{front: 0, length: 0, capcity: k}
	cq.s = make([]int, k)
	return cq
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.length+1 > this.capcity {
		return false
	}
	i := (this.front + this.length) % this.capcity
	this.s[i] = value
	this.length++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.length <= 0 {
		return false
	}
	this.front = (this.front + 1) % this.capcity
	this.length--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.length == 0 {
		return -1
	}
	return this.s[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.length == 0 {
		return -1
	}
	rear := (this.front + this.length - 1) % this.capcity
	return this.s[rear]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.length == this.capcity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end
