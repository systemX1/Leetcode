package Leetcode

import (
    "log"
    "testing"
)

func TestFindRepeatNumber(t *testing.T) {
    fn := findRepeatNumber
    log.Println(fn([]int{2, 3, 1, 0, 2, 5, 3}))
    fn2 := findRepeatNumberNoMod
    log.Println(fn2([]int{0, 1, 2, 0, 4, 5, 6, 7, 8, 9}))

}

func findRepeatNumberNoMod(nums []int) int {
    n := len(nums)
    left, right, mid := 0, n - 1, 0
    for left < right {
        mid = (left + right) >> 1
        if countRange(nums, left, mid) > mid - left + 1 {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func countRange(nums []int, start, end int) int {
    ans := 0
    for _, num := range nums {
        if num >= start && num <= end {
            ans++
        }
    }
    return ans
}

//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatNumber(nums []int) int {
    for i, num := range nums {
        if num == i {
            continue
        }
        for nums[i] != i {
            if nums[i] == nums[nums[i]] {
                return nums[i]
            }
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        }
    }
    return 0
}
//leetcode submit region end(Prohibit modification and deletion)

