package Leetcode

import (
    "log"
    "strings"
    "testing"
)

func TestReplaceSpace(t *testing.T) {
    fn := replaceSpace
    log.Println(fn("We are happy."))
    // Unicode
    for i, ch := range "HELLO, 世界" {
        log.Println(i, ch)
    }
    log.Println()
    // utf-8
    s := "HELLO, 世界"
    for i := 0; i < len(s); i++ {
        log.Println(i, s[i])
    }
}

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
    if len(s) == 0 {
        return ""
    }
    n := len(s)
    for _, re := range s {
        if re == ' ' {
            n++
        }
    }
    var ans strings.Builder
    ans.Grow(n)
    for _, re := range s {
        if re == ' ' {
            ans.WriteString("%20")
        } else {
            ans.WriteRune(re)
        }
    }
    return ans.String()
}
//leetcode submit region end(Prohibit modification and deletion)

