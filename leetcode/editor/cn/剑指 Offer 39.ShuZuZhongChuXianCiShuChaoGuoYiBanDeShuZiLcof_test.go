package Leetcode

import (
    "log"
    "testing"
)

func TestShuZuZhongChuXianCiShuChaoGuoYiBanDeShuZiLcof(t *testing.T) {
    fn := majorityElement
    log.Println(fn([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
    curr, count := nums[0], 0
    for _, num := range nums {
        if num == curr {
            count++
        } else {
            count--
        }
        if count == 0 {
            curr, count = num, 1
        }
    }
    return curr
}
//leetcode submit region end(Prohibit modification and deletion)

