package Leetcode

import (
    "log"
    "testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
    fn := longestCommonSubsequence
    log.Println(fn("abcde", "ace"))
    log.Println(fn("abc", "abc"))
    log.Println(fn("abc", "def"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubsequence(text1 string, text2 string) int {
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }

    m, n := len(text1), len(text2)
    dp := make([][]int, m + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, n + 1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[m][n]
}
//leetcode submit region end(Prohibit modification and deletion)

