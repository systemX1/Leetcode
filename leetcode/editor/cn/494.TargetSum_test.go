package Leetcode

import (
    "log"
    "testing"
)

func TestTargetSum(t *testing.T) {
    fn := findTargetSumWays
    log.Println(fn([]int{1,1,1,1,1}, 3))
    log.Println(fn([]int{1}, 1))
    log.Println(fn([]int{2,107,109,113,127,131,137,3,2,3,5,7,11,13,17,19,23,29,47,53}, 1000))
}

// a = ( Sum + target ) >> 1
//leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if (sum + target) & 1 == 1 || (sum + target) < 0 || target > sum {
        return 0
    }
    target = (sum + target) >> 1

    dp := make([]int, sum + 1)
    dp[0] = 1
    for _, num := range nums {
        for j := target; j >= num ; j-- {
            dp[j] += dp[j-num]
        }
    }
    return dp[target]
}
//leetcode submit region end(Prohibit modification and deletion)

