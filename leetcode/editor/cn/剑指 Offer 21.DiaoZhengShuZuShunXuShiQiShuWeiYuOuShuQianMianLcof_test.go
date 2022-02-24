package Leetcode

import (
    "log"
    "testing"
)

func TestDiaoZhengShuZuShunXuShiQiShuWeiYuOuShuQianMianLcof(t *testing.T) {
    fn := exchange
    log.Println(fn([]int{1,2,3,4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func exchange(nums []int) []int {
    for i, j := 0, len(nums) - 1; i < j; {
        for i < j && nums[i]&1 == 1 {
            i++
        }
        for i < j && nums[j]&1 == 0 {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    return nums
}
//leetcode submit region end(Prohibit modification and deletion)

