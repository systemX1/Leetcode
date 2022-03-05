package Leetcode

import (
    "log"
    "testing"
)

type OTreeNode struct {
    Val int
    Left *OTreeNode
    Right *OTreeNode
}

func TestBuildTree(t *testing.T) {
    fn := buildTree
    log.Println(fn([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func buildTree(preorder []int, inorder []int) *OTreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &OTreeNode{
        Val:   preorder[0],
        Left:  nil,
        Right: nil,
    }
    stk := []*OTreeNode{root}
    idx := 0
    for i := 1; i < len(preorder); i++ {
        top := stk[len(stk)-1]
        if stk[len(stk)-1].Val != inorder[idx] {
            top.Left = &OTreeNode{Val: preorder[i], Left:nil, Right:nil}
            stk = append(stk, top.Left)
        } else {
            for len(stk) > 0 && stk[len(stk)-1].Val == inorder[idx] {
                top = stk[len(stk)-1]
                stk = stk[:len(stk)-1]
                idx++
            }
            top.Right = &OTreeNode{Val: preorder[i], Left:nil, Right:nil}
            stk = append(stk, top.Right)
        }
    }
    return root
}
//leetcode submit region end(Prohibit modification and deletion)

//func buildTree(preorder []int, inorder []int) *OTreeNode {
//    if len(preorder) == 0 {
//        return nil
//    }
//    return build(preorder, inorder)
//}
//
//func build(preorder []int, inorder []int) *OTreeNode {
//    if len(preorder) == 0 {
//        return nil
//    }
//    i := 0
//    for j, v := range inorder {
//        if v == preorder[0] {
//            i = j
//            break
//        }
//    }
//    return &OTreeNode{
//        Val:   preorder[0],
//        Left:  build(preorder[1:len(inorder[:i]) + 1], inorder[:i]),
//        Right: build(preorder[len(inorder[:i]) + 1:], inorder[i+1:]),
//    }
//}

//

//func buildTree(preorder []int, inorder []int) *OTreeNode {
//    if len(preorder) == 0 {
//        return nil
//    }
//    ht := make(map[int]int, len(preorder))
//    for i, v := range inorder {
//        ht[v] = i
//    }
//    return build(preorder, ht, 0, 0, len(preorder) - 1)
//}
//
//func build(preorder []int, ht map[int]int, root, left, right int) *OTreeNode {
//    if left > right {
//        return nil
//    }
//    i := ht[preorder[root]]
//    return &OTreeNode{
//        Val:   preorder[root],
//        Left:  build(preorder, ht, root + 1, left, i - 1),
//        Right: build(preorder, ht, root + i - left + 1, i + 1, right),
//    }
//}

