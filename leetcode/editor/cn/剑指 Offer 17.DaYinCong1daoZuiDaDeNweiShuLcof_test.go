package Leetcode

import (
    "log"
    "testing"
)

func TestPrintNumbers(t *testing.T) {
    fn := printNumbers
    log.Println(fn(1))
    fn2 := printBigNumbers
    var ans [][]string
    ans = append(ans, fn2(1), fn2(2))
    for _, v := range ans {
        log.Println(len(v), v[0], v[len(v)-1])
    }
}

//leetcode submit region begin(Prohibit modification and deletion)
func printNumbers(n int) []int {
    length := pow(10, n) - 1
    ans := make([]int, 0,  length)
    for i := 1; i <= length; i++ {
        ans = append(ans, i)
    }
    return ans
}

func pow(a, b int) int {
    ans := 1
    for b > 0 {
        if b & 1 == 1 {
            ans = ans * a
        }
        b >>= 1
        a = a * a
    }
    return ans
}
//leetcode submit region end(Prohibit modification and deletion)

func printBigNumbers(n int) []string {
    length := pow(10, n) - 1
    ans := make([]string, 0,  length)
    digit := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

    var dfs func(l, n int, curr []byte)
    dfs = func(l, n int, curr []byte) {
        if l >= n {
            ans = append(ans, string(curr))
            return
        }
        start := 0
        if l == 0 {
            start = 1
        }
        for i := start; i < 10; i++ {
            curr = append(curr, digit[i])
            dfs(l + 1, n, curr)
            curr = curr[:len(curr)-1]
        }

    }
    for i := 1; i <= n; i++ {
        dfs(0, i, make([]byte, 0, i))
    }
    return ans
}