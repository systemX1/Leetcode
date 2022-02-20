/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */
package leetcode

// @lc code=start
func singleNonDuplicate(nums []int) int {
	l := len(nums)
	left, right, mid := 0, l-1, 0
	for left < right {
		mid = (left + right) >> 1
		mid -= mid & 1
		if nums[mid] == nums[mid+1] {
			left += 2
		} else {
			right = mid
		}
	}
	return nums[left]
}

// @lc code=end
