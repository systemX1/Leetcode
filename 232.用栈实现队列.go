/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */
package leetcode

// @lc code=start
type MyQueue struct {
	s1     []int
	s2     []int
	front1 int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: []int{},
		s2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	if len(this.s1) == 0 {
		this.front1 = x
	}
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.s2) == 0 {
		for len(this.s1) != 0 {
			top1 := this.s1[len(this.s1)-1]
			this.s1 = this.s1[:len(this.s1)-1]
			this.s2 = append(this.s2, top1)
		}
	}
	top2 := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return top2
}

func (this *MyQueue) Peek() int {
	if len(this.s2) != 0 {
		return this.s2[len(this.s2)-1]
	}
	return this.front1
}

func (this *MyQueue) Empty() bool {
	return len(this.s1)+len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
