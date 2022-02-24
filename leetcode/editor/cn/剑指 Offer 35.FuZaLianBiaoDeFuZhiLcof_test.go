package Leetcode

import (
    "fmt"
    "log"

    "strings"
    "testing"
)

func TestCopyRandomList(t *testing.T) {
    //fn := copyRandomList
    list := constructList([][]int{
        {7,-1},{13,0},{11,4},{10,2},{1,0},
    })
    log.Println(printList(list))
    log.Println(printList(copyRandomList(list) ) )

    list2 := constructList([][]int{
        {7,1},{13,0},
    })
    log.Println(printList(list2))
    log.Println(printList(copyRandomList(list2) ) )
}

type Node struct {
    Val int
    Next *Node
    Random *Node
}

func constructList(data [][]int) *Node {
    if len(data) == 0 {
        return nil
    }
    dummy := &Node{}
    ht := make(map[int]*Node, len(data))
    ht[-1] = nil
    for i, p := 0, dummy; i < len(data); i, p = i+1, p.Next {
        p.Next = &Node{Val: data[i][0]}
        ht[i] = p.Next
    }
    for i, p := 0, dummy.Next; p != nil; i, p = i+1, p.Next {
        if v, ok := ht[data[i][1]]; ok {
            p.Random = v
        }
    }
    return dummy.Next
}

func printList(head *Node) string {
    var sb strings.Builder
    sb.WriteString("[ ")
    for p := head; p != nil; p = p.Next {
        randomVal := -1
        if p.Random != nil {
            randomVal = p.Random.Val
        }
        sb.WriteString(fmt.Sprintf("%v,%v ", p.Val, randomVal))
    }
    sb.WriteRune(']')
    return sb.String()
}
//leetcode submit region begin(Prohibit modification and deletion)
func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    for p := head; p != nil; p = p.Next.Next {
        p.Next = &Node{Val: p.Val, Next: p.Next}
    }
    for p, p2 := head, head.Next; p != nil; p = p.Next.Next {
        p2 = p.Next
        if p.Random != nil {
            p2.Random = p.Random.Next
        }
    }
    dummy2 := &Node{Next: head.Next}
    for p, p2 := head, dummy2; p != nil; p = p.Next {
        p2.Next = p.Next
        p2 = p2.Next
        p.Next = p2.Next
    }
    return dummy2.Next
}
//leetcode submit region end(Prohibit modification and deletion)

// hashtable
//func copyRandomList(head *Node) *Node {
//    if head == nil {
//        return nil
//    }
//    ht := map[*Node]*Node{}
//    dummy2 := &Node{-1, nil, nil}
//    p2 := dummy2
//    for p := head; p != nil; p = p.Next {
//        p2.Next = &Node{Val: p.Val, Next: nil, Random: nil}
//        p2 = p2.Next
//        ht[p] = p2
//    }
//    for p, p2 := head, dummy2.Next; p2 != nil; p, p2 = p.Next, p2.Next {
//        p2.Random = ht[p.Random]
//    }
//    return dummy2.Next
//}
