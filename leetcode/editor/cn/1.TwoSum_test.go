package Leetcode

import (
    "log"
    "testing"
)

func TestTwoSum(t *testing.T) {
    fn := twoSum
    log.Println(fn([]int{2,7,11,15}, 9))
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
    ht := map[int]int{}
    for i, num := range nums {
        if n, ok := ht[target-num]; ok {
            return []int{n, i}
        }
        ht[num] = i
    }
    return nil
}
//leetcode submit region end(Prohibit modification and deletion)

