/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */
package leetcode

// @lc code=start
type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	l := len(this.queue)
	this.queue = append(this.queue, x)
	for i := 0; i < l; i++ {
		front := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, front)
	}
}

func (this *MyStack) Pop() int {
	front := this.queue[0]
	this.queue = this.queue[1:]
	return front
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
