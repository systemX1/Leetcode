package Leetcode

import (
    "log"
    "testing"
)

func TestThreeStepsProblemLcci(t *testing.T) {
    fn := waysToStep
    log.Println(fn(3), fn(5))
}

//leetcode submit region begin(Prohibit modification and deletion)
func waysToStep(n int) int {
    switch n {
    case 3:
        return 4
    case 2:
        return 2
    case 1:
        return 1
    }
    v1, v2, v3, v4 := 1, 2, 4, 7
    for i := 4; i <= n; i++ {
        v4 = (v1 + v2 + v3) % (1e9+7)
        v1, v2, v3 = v2, v3, v4
    }
    return v4
}
//leetcode submit region end(Prohibit modification and deletion)

