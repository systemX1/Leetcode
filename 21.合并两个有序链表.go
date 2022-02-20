/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{-1, nil}
	curr, p1, p2 := preHead, list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			curr.Next = p1
			curr = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			curr = p2
			p2 = p2.Next
		}
	}
	if p1 == nil {
		curr.Next = p2
	}
	if p2 == nil {
		curr.Next = p1
	}
	return preHead.Next
}

// @lc code=end
