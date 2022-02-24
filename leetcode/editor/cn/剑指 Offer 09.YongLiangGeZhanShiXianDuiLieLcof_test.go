package Leetcode

import (
    "testing"
)

func TestYongLiangGeZhanShiXianDuiLieLcof(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
type CQueue struct {
    stk1 []int
    stk2 []int
}

func Constructor() CQueue {
    return CQueue{
        stk1: []int{},
        stk2: []int{},
    }
}

func (q *CQueue) AppendTail(value int)  {
    q.stk1 = append(q.stk1, value)
}

func (q *CQueue) DeleteHead() int {
    if len(q.stk1) == 0 && len(q.stk2) == 0 {
        return -1
    }
    if len(q.stk2) == 0 {
        for len(q.stk1) > 0 {
            q.stk2 = append(q.stk2, q.stk1[len(q.stk1)-1])
            q.stk1 = q.stk1[:len(q.stk1)-1]
        }
    }
    top := q.stk2[len(q.stk2)-1]
    q.stk2 = q.stk2[:len(q.stk2)-1]
    return top
}

//leetcode submit region end(Prohibit modification and deletion)
/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
