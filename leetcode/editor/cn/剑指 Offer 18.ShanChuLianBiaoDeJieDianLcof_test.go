package Leetcode

import (
    "log"
    "testing"
)

func TestDeleteNode(t *testing.T) {
    fn := deleteNode
    log.Println(fn(nil, 0))
}

type ListNode struct {
     Val  int
     Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }
    dummy := &ListNode{Val: -1, Next: head}
    for slow, fast := dummy, head; fast != nil; slow, fast = slow.Next, fast.Next {
        if fast.Val == val {
            slow.Next = slow.Next.Next
        }
    }
    return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)

