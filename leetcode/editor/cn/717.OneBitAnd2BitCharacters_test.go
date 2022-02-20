package Leetcode

import (
    "log"
    "testing"
)

func TestOneBitAnd2BitCharacters(t *testing.T) {
    fn := oneBitAnd2BitCharacters
    log.Println(fn([]int{1,0,0}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isOneBitCharacter(bits []int) bool {
    i := 0
    for i < len(bits) {
        if bits[i] == 0 {
            i++
        } else if bits[i] == 1 {
            i += 2
        }
    }
    return false
}
//leetcode submit region end(Prohibit modification and deletion)

