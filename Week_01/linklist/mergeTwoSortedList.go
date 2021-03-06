package linklist

type ListNode struct {
	Val int
	Next *ListNode
 }
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    prevHead := new(ListNode)
    prev := prevHead
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
            prev.Next = l2
            l2 = l2.Next
        } else {
            prev.Next = l1
            l1 = l1.Next
        }
        prev = prev.Next
    }
    if l1 == nil {
        prev.Next = l2
    }
    if l2 == nil {
        prev.Next = l1
    }
    return prevHead.Next
}