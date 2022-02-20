/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
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
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	preSm, preLr := &ListNode{}, &ListNode{}
	currSm, currLr := preSm, preLr
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val < x {
			currSm.Next = curr
			currSm = currSm.Next
		} else {
			currLr.Next = curr
			currLr = currLr.Next
		}
	}
	currLr.Next = nil
	currSm.Next = preLr.Next
	return preSm.Next
}

// @lc code=end
