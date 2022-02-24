package Leetcode

import (
    "log"
    "testing"
)

func TestOneBitAnd2BitCharacters(t *testing.T) {
    fn := isOneBitCharacter
    log.Println(fn([]int{1,0,0}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isOneBitCharacter(bits []int) bool {
    i, n := 0, len(bits)
    for i <= n - 2 {
        i += bits[i] + 1
    }
    return i == n - 1
}
//leetcode submit region end(Prohibit modification and deletion)

