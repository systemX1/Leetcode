/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */
package leetcode

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	stack, ret := []int{}, make([]int, l)
	for i := 0; i < l; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[top] = i - top
		}
		stack = append(stack, i)
	}
	return ret
}

// @lc code=end
