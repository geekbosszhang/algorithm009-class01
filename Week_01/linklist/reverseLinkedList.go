package linklist

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var prev *ListNode = nil
    cur := head
    for cur != nil {
        nextTmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = nextTmp
    }
    return prev
}