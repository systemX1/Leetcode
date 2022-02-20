// @before-stub-for-debug-begin
#include <vector>
#include <string>
#include <unordered_set>
#include "commoncppproblem3.h"

using namespace std;
// @before-stub-for-debug-end

/*
 * @lc app=leetcode.cn id=3 lang=cpp
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        if(s.empty())   return 0;
        int ret = 0, right = -1;
        unordered_set<char> uomap;
        for (int left = 0; left < s.size(); left++) {
            if(left)
                uomap.erase(s[left - 1]);
            while(!uomap.count(s[right + 1]) && right + 1 < s.size()) {
                uomap.insert(s[right + 1]);
                right++;
            }
            ret = max(ret, right - left + 1);
        }
        
        return ret;
    }
};
// @lc code=end

