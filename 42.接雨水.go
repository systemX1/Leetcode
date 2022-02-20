/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package leetcode

// @lc code=start
func trap(height []int) int {
	l := len(height)
	if l <= 2 {
		return 0
	}
	stack := make([]int, 0)
	ret := 0
	for i := 0; i < l; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			width := i - stack[len(stack)-1] - 1
			boundedHeight := min(height[i], height[stack[len(stack)-1]]) - height[top]
			ret += width * boundedHeight
		}
		stack = append(stack, i)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
