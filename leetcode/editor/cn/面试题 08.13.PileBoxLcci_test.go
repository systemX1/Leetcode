package Leetcode

import (
    "log"
    "sort"
    "testing"
)

func TestPileBoxLcci(t *testing.T) {
    fn := pileBox
    log.Println(fn([][]int{
        {1, 1, 1}, {2, 2, 2}, {3, 3, 3},
    }))
    log.Println(fn([][]int{
        {1, 1, 1}, {2, 3, 4}, {2, 6, 7}, {3, 4, 5},
    }))
}

//leetcode submit region begin(Prohibit modification and deletion)
func pileBox(box [][]int) int {
    n := len(box)
    if n <= 1 {
        return n
    }
    sort.Slice(box, func(i, j int) bool {
        a, b := box[i], box[j]
        return a[2] > b[2]
    })
    dp := make([]int, n)
    ans := 0
    for i := 0; i < n; i++ {
        dp[i] = box[i][2]
        for j := 0; j < i; j++ {
            if box[i][0] < box[j][0] && box[i][1] < box[j][1] && box[i][2] < box[j][2] && dp[j] + box[i][2] > dp[i] {
                dp[i] = dp[j] + box[i][2]
            }
        }
        if dp[i] > ans {
            ans = dp[i]
        }
    }
    return ans
}
//leetcode submit region end(Prohibit modification and deletion)

