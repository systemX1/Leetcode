/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package leetcode

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	deque := []int{}
	ans := make([]int, 0, l-k+1)

	for i := 0; i < k; i++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}

	ans = append(ans, nums[deque[0]])
	for i := k; i < l; i++ {
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		ans = append(ans, nums[deque[0]])
	}

	return ans
}

// @lc code=end
