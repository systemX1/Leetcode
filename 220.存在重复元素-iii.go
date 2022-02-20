/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */
package leetcode

// @lc code=start
func getID(x, size int) int {
	if x >= 0 {
		return x / size
	}
	return (x+1)/size - 1
}

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	ht := map[int]int{}
	for i, num := range nums {
		id := getID(num, t+1)
		if _, ok := ht[id]; ok {
			return true
		}
		if v, ok := ht[id-1]; ok && abs(num-v) <= t {
			return true
		}
		if v, ok := ht[id+1]; ok && abs(num-v) <= t {
			return true
		}
		ht[id] = num
		if i >= k {
			delete(ht, getID(nums[i-k], t+1))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
