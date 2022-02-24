package Leetcode

import (
    "log"
    "math"
    "testing"
)

func TestTriangle(t *testing.T) {
    fn := minimumTotal
    log.Println(fn([][]int{
        {2},
        {3,4},
        {6,5,7},
        {4,1,8,3},
    }))
    log.Println(fn([][]int{
       {-10},
    }))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    m := len(triangle)
    dp := make([]int, m)
    for i := 0; i < m; i++ {
        dp[i] = math.MaxInt64
    }
    dp[0] = triangle[0][0]
    for i := 1; i < m; i++ {
        dp[i] = dp[i-1] + triangle[i][i]
        for j := i - 1; j >= 1; j-- {
            dp[j] = min(dp[j] + triangle[i][j], dp[j-1] + triangle[i][j])
        }
        dp[0] = dp[0] + triangle[i][0]
    }
    ans := math.MaxInt64
    for _, v := range dp {
        ans = min(ans, v)
    }
    return ans
}
//leetcode submit region end(Prohibit modification and deletion)

