/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	ht := map[int]int{}
	for i, num := range nums {
		if n, ok := ht[target-num]; ok {
			return []int{n, i}
		}
		ht[num] = i
	}
	return nil
}

// @lc code=end
