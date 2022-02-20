/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{-1, head}
	p := dummy
	for left > 1 {
		p = p.Next
		left--
		right--
	}
	q := p.Next
	for right > 1 {
		tmp := p.Next
		p.Next = q.Next
		q.Next = q.Next.Next
		p.Next.Next = tmp
		right--
	}

	return dummy.Next
}

// @lc code=end
