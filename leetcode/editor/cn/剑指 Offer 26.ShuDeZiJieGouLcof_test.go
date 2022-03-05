package Leetcode

import (
    "log"
    "testing"
)

func TestShuDeZiJieGouLcof(t *testing.T) {
    fn := isSubStructure
    log.Println(fn(nil, nil))
}

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type OTreeNode struct {
 *     Val int
 *     Left *OTreeNode
 *     Right *OTreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if B == nil || A == nil {
        return false
    }


}
//leetcode submit region end(Prohibit modification and deletion)

