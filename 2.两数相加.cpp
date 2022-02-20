// @before-stub-for-debug-begin
#include <vector>
#include <string>
#include "commoncppproblem2.h"

using namespace std;
// @before-stub-for-debug-end

/*
 * @lc app=leetcode.cn id=2 lang=cpp
 *
 * [2] 两数相加 [2,4,3]\n[5,6,4]
 */
// struct ListNode {
//     int val;
//     ListNode *next;
//     ListNode() : val(0), next(nullptr) {}
//     ListNode(int x) : val(x), next(nullptr) {}
//     ListNode(int x, ListNode *next) : val(x), next(next) {}
// };
// @lc code=start
/**
 * Definition for singly-linked list.
 */

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        if(!l1) return l2;
        if(!l2) return l1;
        ListNode *curr = new ListNode();
        ListNode *dummy = curr;
        int carry = 0;
        while (l1 || l2)
        {
            int x = 0, y = 0;
            if(l1)
                x = l1->val;
            if(l2)
                y = l2->val;
            curr->next = new ListNode( (x + y + carry) % 10);
            curr = curr->next;
            carry = (x + y + carry) / 10;

            if(l1)
                l1 = l1->next;
            if(l2)
                l2 = l2->next;
        }

        if(carry)
            curr->next = new ListNode(carry);
        return dummy->next;
    }
};
// @lc code=end

