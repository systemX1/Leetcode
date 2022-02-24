package Leetcode

import (
    "log"
    "testing"
)

func TestMinimumPathSum(t *testing.T) {
    fn := minPathSum
    log.Println(fn([][]int{
        {1,3,1},
        {1,5,1},
        {4,2,1},
    }))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    dp[0][0] = grid[0][0]

    for i := 1; i < m; i++ {
        dp[i][0] = dp[i-1][0] + grid[i][0]
    }
    for j := 1; j < n; j++ {
        dp[0][j] = dp[0][j-1] + grid[0][j]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = Omin(dp[i-1][j] + grid[i][j], dp[i][j-1] + grid[i][j])
        }
    }
    return dp[m-1][n-1]
}

func Omin(a, b int) int {
    if a < b {
        return a
    }
    return b
}
//leetcode submit region end(Prohibit modification and deletion)

