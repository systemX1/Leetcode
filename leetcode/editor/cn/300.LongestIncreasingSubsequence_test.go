package Leetcode

import (
    "log"
    "testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
    fn := lengthOfLIS
    log.Println(fn([]int{10,9,2,5,3,7,101,18}))
    log.Println(fn([]int{0,1,0,3,2,3}))
    log.Println(fn([]int{0}))
    log.Println(fn([]int{7,7,7,7}))
    log.Println(fn([]int{3,5,6,2,5,4,19,5,6,7,12}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    ans := 0
    for i := 0; i < n; i++ {
        l, r := 0, ans
        for l < r {
            mid := (l + r) >> 1
            if dp[mid] < nums[i] {
                l = mid + 1
            } else {
                r = mid
            }
        }
        dp[l] = nums[i]
        if l == ans {
            ans = l + 1
        }
    }
    return ans
}
//leetcode submit region end(Prohibit modification and deletion)

// dp O(n^2)
//func lengthOfLIS(nums []int) int {
//    n := len(nums)
//    dp := make([]int, n)
//    dp[0] = 1
//    res := 1
//    for i := 1; i < n; i++ {
//        dp[i] = 1
//        for j := 0; j < i; j++ {
//            if nums[j] < nums[i] && dp[i] < dp[j] + 1 {
//                dp[i] = dp[j] + 1
//            }
//            if dp[i] > res {
//                res = dp[i]
//            }
//        }
//    }
//    return res
//}

