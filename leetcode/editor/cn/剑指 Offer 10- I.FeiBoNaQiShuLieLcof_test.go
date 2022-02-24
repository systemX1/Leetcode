package Leetcode

import (
    "log"
    "testing"
)

func TestFib(t *testing.T) {
    fn := fib
    log.Println(fn(2), fn(5), fn(95))
    log.Println(int(1407059028) % (1e9 + 7))

}

//leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
    if n < 2 {
        return n
    }
    f, s, t := 0, 1, 1
    for i := 2; i <= n; i++ {
        t = (f + s) % (1e9 + 7)
        f, s = s, t
    }
    return t
}
//leetcode submit region end(Prohibit modification and deletion)

