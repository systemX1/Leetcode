package Leetcode

import (
    "log"
    "testing"
)

func TestPowxN(t *testing.T) {
    fn := myPow
    log.Println(fn(2.1, 3), fn(2.1, -2), fn(-2.0, 3), fn(-2.0, -3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
    ans := 1.0
    isNegative := false
    if n < 0 {
        isNegative = true
        n = -n
    }
    for n > 0 {
        if n & 1 == 1 {
            ans = ans * x
        }
        n >>= 1
        x *= x
    }
    if isNegative {
        return 1.0 / ans
    }
    return ans
}
//leetcode submit region end(Prohibit modification and deletion)

