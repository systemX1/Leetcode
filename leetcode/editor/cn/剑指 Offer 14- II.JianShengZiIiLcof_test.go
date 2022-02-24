package Leetcode

import (
    "log"
    "testing"
)

func TestCuttingRope2(t *testing.T) {
    //fn := cuttingRope
    //log.Println(fn(10))
    fastPower := func(x, a, p int) int {
        ans := 1
        x %= p
        for a > 0 {
            if a & 1 == 1 {
                ans = (ans * x) % p
            }
            x = (x * x) % p
            a >>= 1
        }
        return ans
    }
    log.Println(fastPower(3, 5 ,7))
    log.Println(3439254224 % int(1e9+7))
    log.Println(9 * 4,  9 * 4 % int(1e9+7))
    fn := cuttingRope
    log.Println(fn(10))
}

//leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
    if n < 4 {
        return n - 1
    }
    fastPow := func(a, b, p int) int {
        ans := 1
        for b > 0 {
            if b & 1 == 1 {
                ans = (ans * a) % p
            }
            b >>= 1
            a = (a * a) % p
        }
        return ans
    }
    switch n % 3 {
    case 2:
        return fastPow(3, n/3, 1e9+7) * 2 % int(1e9+7)
    case 1:
        return fastPow(3, n/3 - 1, 1e9+7) * 4 % int(1e9+7)
    default:
        return fastPow(3, n/3, 1e9+7) % int(1e9+7)
    }

}
//leetcode submit region end(Prohibit modification and deletion)

