/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */
package leetcode

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	l := len(pushed)
	stack := []int{}
	i, j := 0, 0
	for i < l {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
		i++
	}
	return len(stack) == 0
}

// @lc code=end
