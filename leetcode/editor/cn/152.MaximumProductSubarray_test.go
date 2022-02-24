package Leetcode

import (
    "log"
    "testing"
)

func TestMaximumProductSubarray(t *testing.T) {
    fn := maxProduct
    log.Println(fn([]int{2,3,-2,4}), fn([]int{-4,-3,-2}))
    log.Println(SumO(1, 2, 3, 4))
    s := []int{1, 2, 3, 4}
    log.Println(SumO(s...))

}

func SumO(nums ...int) (ans int) {
    for _, r := range nums{
        ans += r
    }
    return
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
    dpMin, dpMax, ans := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        tmpMin, tmpMax := dpMin, dpMax
        if nums[i] >= 0 {
            dpMax = max(tmpMax * nums[i], nums[i])
            dpMin = min(tmpMin * nums[i], nums[i])
        } else {
            dpMax = max(tmpMin * nums[i], nums[i])
            dpMin = min(tmpMax * nums[i], nums[i])
        }
        ans = max(dpMax, ans)
    }
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
//leetcode submit region end(Prohibit modification and deletion)

