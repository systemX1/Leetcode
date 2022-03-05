package Leetcode

import (
    "log"
    "sort"
    "testing"
)

func TestRussianDollEnvelopes(t *testing.T) {
    fn := maxEnvelopes
    log.Println(fn([][]int{
        {5,4},{6,4},{6,7},{2,3},
    }))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxEnvelopes(envelopes [][]int) int {
    n := len(envelopes)
    sort.Slice(envelopes, func(i, j int) bool {
        a, b := envelopes[i], envelopes[j]
        return a[1] < b[1] || (a[1] == b[1] && a[0] > b[0])
    })
    dp := make([]int, n)
    res := 0
    for _, e := range envelopes {
        l, r := 0, res
        for l < r {
            mid := (l + r) >> 1
            if dp[mid] < e[0] {
                l = mid + 1
            } else {
                r = mid
            }
        }
        dp[l] = e[0]
        if l == res {
            res = l + 1
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)

