package Leetcode

import (
    "log"
    "testing"
)

func TestReversePrint(t *testing.T) {
    fn := reversePrint
    log.Println(fn(nil))
}

type OListNode struct {
    Val int
    Next *OListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *OListNode) []int {
    var ans, stk []int
    if head == nil {
        return ans
    }
    for p := head; p != nil; p = p.Next {
        stk = append(stk, p.Val)
    }
    for len(stk) > 0 {
        ans = append(ans, stk[len(stk)-1])
        stk = stk[:len(stk)-1]
    }
    return ans
}
//leetcode submit region end(Prohibit modification and deletion)

