/*
 * @lc app=leetcode.cn id=7 lang=cpp
 *
 * [7] 整数反转
 */

// @lc code=start
class Solution {
public:
    int reverse(int x) {
        int ret = 0;
        while (x) {
            if (ret > __INT_MAX__ / 10 || ret < -__INT_MAX__ / 10 )
                return 0;
            ret = ret * 10 + x % 10;
            x /= 10;
        }
        return ret;
    }
};
// @lc code=end
