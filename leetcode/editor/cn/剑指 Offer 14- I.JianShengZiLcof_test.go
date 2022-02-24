package Leetcode

import (
    "log"
    "testing"
)

func TestCuttingRope(t *testing.T) {
    fn := cuttingRope
    log.Println(fn(10))
}

//leetcode submit region begin(Prohibit modification and deletion)
//func cuttingRope(n int) int {
//    if n < 4 {
//        return n - 1
//    }
//    switch n % 3 {
//    case 2:
//        return int(math.Pow(3, float64(n / 3)) * 2)
//    case 1:
//        return int(math.Pow(3, float64(n / 3 - 1)) * 4)
//    default:
//        return int(math.Pow(3, float64(n / 3)) )
//    }
//}
//leetcode submit region end(Prohibit modification and deletion)

