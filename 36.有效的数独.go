/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */
package leetcode

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	row, col, ht := [9][9]int{}, [9][9]int{}, [9][9]int{}
	for i, r := range board {
		for j, ch := range r {
			if ch == '.' {
				continue
			}
			elem := ch - '1'
			idx := getHtIdx(i, j)
			if row[i][elem] != 0 || col[j][elem] != 0 || ht[idx][elem] != 0 {
				return false
			}
			row[i][elem], col[j][elem], ht[idx][elem] = 1, 1, 1
		}
	}
	return true
}

func getHtIdx(i, j int) int {
	return (i/3)*3 + j/3
}

// @lc code=end
