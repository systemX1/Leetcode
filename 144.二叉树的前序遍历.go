/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) (ret []int) {
	var preorder func(curr *TreeNode)
	preorder = func(curr *TreeNode) {
		if curr == nil {
			return
		}
		ret = append(ret, curr.Val)
		preorder(curr.Left)
		preorder(curr.Right)
	}
	preorder(root)
	return ret
}

// @lc code=end
