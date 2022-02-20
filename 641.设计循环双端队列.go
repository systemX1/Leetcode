/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */
package leetcode

// @lc code=start
type MyCircularDeque struct {
	s       []int
	front   int
	length  int
	capcity int
}

func Constructor(k int) MyCircularDeque {
	cd := MyCircularDeque{front: 0, length: 0, capcity: k}
	cd.s = make([]int, k)
	return cd
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length+1 > this.capcity {
		return false
	}
	this.front = (this.front + this.capcity - 1) % this.capcity
	this.s[this.front] = value
	this.length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length+1 > this.capcity {
		return false
	}
	i := (this.front + this.length) % this.capcity
	this.s[i] = value
	this.length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.length <= 0 {
		return false
	}
	this.front = (this.front + 1) % this.capcity
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.length <= 0 {
		return false
	}
	this.length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.length == 0 {
		return -1
	}
	return this.s[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.length == 0 {
		return -1
	}
	rear := (this.front + this.length - 1) % this.capcity
	return this.s[rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.capcity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end
