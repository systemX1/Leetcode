/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */
package leetcode

// @lc code=start
func findPeakElement(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	left, right, mid := 0, l-1, 0
	for left < right {
		mid = (left + right) >> 1
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end
