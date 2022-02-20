package leetcode

import (
	"bytes"
)

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	br := func() int {
		if numRows < len(s) {
			return numRows
		}
		return len(s)
	}()
	bufS := make([]bytes.Buffer, br, br)
	r := 0
	flag := -1
	for i := 0; i < len(s); i++ {
		bufS[r].WriteByte(s[i])
		if r == numRows-1 || r == 0 {
			flag = -flag
		}
		r += flag
	}
	var buf bytes.Buffer
	for _, bs := range bufS {
		buf.Write(bs.Bytes())
	}
	return buf.String()
}

// @lc code=end
