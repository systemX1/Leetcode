package Leetcode

import (
    "log"
    "testing"
)

func TestHouseRobber(t *testing.T) {
    fn := rob
    log.Println(fn([]int{2,7,9,3,1}))
    log.Println(fn([]int{1,2,3,1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    dp := make([]int, n)
    dp[0], dp[1] = nums[0], max(nums[0], nums[1])
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-2] + nums[i], dp[i-1])
    }
    return dp[n-1]
}
//leetcode submit region end(Prohibit modification and deletion)

