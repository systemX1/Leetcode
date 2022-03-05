package Leetcode

import (
    "log"
    "strconv"
    "strings"
    "testing"
)

func TestSerializeAndDeserializeBinaryTree(t *testing.T) {
    ser := Constructor()
    deser := Constructor()
    data := ser.serialize(nil)
    ans := deser.deserialize(data)
    log.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
    var sb strings.Builder
    stk := []*TreeNode{root}
    for len(stk) > 0 {
        node := stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        if node != nil {
            sb.WriteString(strconv.Itoa(node.Val))
            sb.WriteString(",")
            stk = append(stk, node.Left)
            stk = append(stk, node.Right)
        } else {
            sb.WriteString("null,")
        }
    }
    return sb.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
    sp := strings.Split(data, ",")
    if len(sp) == 0 || sp[0] == "null" {
        return nil
    }
    val, _ := strconv.Atoi(sp[0])
    sp = sp[1:]
    root := &TreeNode{Val: val}
    stk := []*TreeNode{root}
    for len(sp) > 0 && len(stk) > 0{
        node := stk[len(stk)-1]
        stk = stk[:len(stk)-1]

        if sp[0] != "null" {
            val, _ := strconv.Atoi(sp[0])
            node.Left = &TreeNode{Val: val}
        } else {
            sp = sp[1:]
        }

        if sp[0] != "null" {
            val, _ := strconv.Atoi(sp[0])
            node.Right = &TreeNode{Val: val}
        } else {
            sp = sp[1:]
        }
    }
    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)

