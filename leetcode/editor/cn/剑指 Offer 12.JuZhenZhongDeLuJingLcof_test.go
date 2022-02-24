package Leetcode

import (
    "log"
    "testing"
)

func TestExist(t *testing.T) {
    fn := exist
    log.Println(fn([][]byte{
        {'A','B','C','E'},
        {'S','F','C','S'},
        {'A','D','E','E'},
    }, "A"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
    var dfs func(i, j, n int) bool
    dfs = func(i, j, n int) bool {
        if n >= len(word) {
            return true
        }
        if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[n] {
            return false
        }
        board[i][j] = '#'
        defer func() {board[i][j] = word[n]}()
        return dfs(i-1, j, n+1) || dfs(i+1, j, n+1) || dfs(i, j-1, n+1) || dfs(i, j+1, n+1)
    }
    for i, row := range board {
        for j := range row {
            if dfs(i, j, 0) {
                return true
            }
        }
    }
    return false
}
//leetcode submit region end(Prohibit modification and deletion)

