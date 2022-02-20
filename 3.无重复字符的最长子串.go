/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	ht := map[byte]int{}
	l, left, ans := len(s), 0, 0
	for i := 0; i < l; i++ {
		if v, ok := ht[s[i]]; ok {
			for j := left; j <= v; j++ {
				delete(ht, s[j])
			}
			left = v + 1
		}
		ht[s[i]] = i
		if i-left+1 > ans {
			ans = i - left + 1
		}
	}
	return ans
}

// @lc code=end
