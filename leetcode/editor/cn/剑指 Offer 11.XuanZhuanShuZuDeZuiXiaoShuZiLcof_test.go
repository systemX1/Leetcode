package Leetcode

import (
    "log"
    "testing"
)

func TestMinArray(t *testing.T) {
    fn := minArray
    log.Println(fn([]int{3,4,5,1,2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minArray(numbers []int) int {
    left, right, mid := 0, len(numbers) - 1, 0
    for left < right {
        mid = left + (right - left) >> 1
        if numbers[mid] > numbers[right] {
            left = mid + 1
        } else if numbers[mid] < numbers[right] {
            right = mid
        } else {
            right--
        }
    }
    return numbers[left]
}
//leetcode submit region end(Prohibit modification and deletion)

