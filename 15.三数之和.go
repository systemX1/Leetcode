package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	rets := make([][]int, 0, 0)

	for first := 0; first < l-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		if nums[first] > 0 {
			break
		}
		third := l - 1
		for second := first + 1; second < l-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for nums[first]+nums[second]+nums[third] > 0 &&
				second < third {
				third--
			}
			if second == third {
				break
			}

			if nums[first]+nums[second]+nums[third] == 0 {
				rets = append(rets,
					[]int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return rets
}

// @lc code=end
