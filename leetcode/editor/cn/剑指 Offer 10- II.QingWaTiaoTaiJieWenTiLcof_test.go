package Leetcode

import (
    "log"
    "testing"
)

func TestNumWays(t *testing.T) {
    fn := numWays
    log.Println(fn(7))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numWays(n int) int {
    if n < 2 {
        return 1
    }
    f, s, t := 1, 1, 2
    for i := 2; i <= n; i++ {
        t = (f + s) % (1e9 + 7)
        f, s = s, t
    }
    return t
}
//leetcode submit region end(Prohibit modification and deletion)

