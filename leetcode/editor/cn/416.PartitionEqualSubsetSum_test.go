package Leetcode

import (
    "log"
    "testing"
)

func TestPartitionEqualSubsetSum(t *testing.T) {
    fn := canPartition
    log.Println(fn([]int{1,5,11,5}))
    log.Println(fn([]int{1,2,3,5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum & 1 == 1 {
        return false
    }
    sum >>= 1
    dp := make([]int, sum + 1)
    dp[0] = 1
    for _, num := range nums {
        for j := sum; j >= num ; j-- {
            if dp[j-num] == 1 {
                dp[j] = 1
                if j == sum {
                    return true
                }
            }
        }
    }
    return dp[sum] == 1
}
//leetcode submit region end(Prohibit modification and deletion)

