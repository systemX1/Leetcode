package Leetcode

import (
    "log"
    "testing"
)

func TestFindNumberIn2DArray(t *testing.T) {
    fn := findNumberIn2DArray
    log.Println(fn([][]int{
       {1, 4, 7, 11, 15},
       {2, 5, 8, 12, 19},
       {10, 13, 14, 17, 24},
       {18, 21, 23, 26, 30},
    }, 10))
    log.Println(fn([][]int{
        {1, 4, 7, 11, 15},
        {2, 5, 8, 12, 19},
        {10, 13, 14, 17, 24},
        {18, 21, 23, 26, 30},
    }, 20))
    log.Println(fn([][]int{
       {},
    }, 5))
    log.Println(fn([][]int{
        {-1, 3},
    }, 5))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    n, m := len(matrix), len(matrix[0])
    h, v := 0, m-1
    for (matrix[h][v] > target && v > 0) || (matrix[h][v] < target && h < n - 1) {
        for matrix[h][v] > target && v > 0 {
            v--
        }
        for matrix[h][v] < target && h < n - 1 {
            h++
        }
    }
    return matrix[h][v] == target
}
//leetcode submit region end(Prohibit modification and deletion)

