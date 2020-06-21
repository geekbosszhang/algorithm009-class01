学习笔记

*链表常见的习题*

1. 合并两个有序列表
   思路：当我们处理两个链表中值较小的节点合并到已经合并的链表中后，剩余的节点依然是有序的，因此是一个递归的过程
   ```
   func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
        if l1 == nil {
            return l2
        }
        if l2 == nil {
            return l1
        }
        if l1.Val < l2.Val {
            l1.Next = mergeTwoLists(l1.Next, l2)
            return l1
        } else {
            l2.Next = mergeTwoLists(l1, l2.Next)
            return l2
        }
    }
   ```
2. 反转链表 
   思路：比如 a->b->...->h->i->j... 前面的部分已经调节好，现在处理到i，调换h<-i
   需要保持i(cur),i的prev，i的next节点j, 以防止节点断裂

    ```
    if head == nil || head.next == nil {
        return head
    }
    prev := nil
    cur := head
    for (cur!=nil) {
        next := cur.next
        cur.next = prev
        prev = cur
        cur = next
    }
    return prev
    ```                                                                                                                                                                                                     
3. 两两交换链表
   思路：head -> 1 -> 2 -> 3 -> 4
   p - head, cur - 1,
   step 1. p.next = cur.next  (head -> 2)
   step 2. cur.next = cur.next.next (1 -> 3)
   step 3. p.next.next = cur (2 -> 1)
   step 4. p = cur
   step 5. cur = cur.next
   ```
   func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    tmp := &ListNode{0, head}
    prev := tmp
    cur := prev.Next
    for cur != nil && cur.Next != nil {
        prev.Next = cur.Next
        cur.Next = cur.Next.Next
        prev.Next.Next = cur

        prev = cur
        cur = cur.Next
    }
    return tmp.Next

}
   ```

   思路2：用递归
   // 1 -> 2 -> 3 -> 4
   // 2 -> 1 -> 4 -> 3
   // (2 tail) -> (1 head) -> swapPairs(3 -> 4)
   // rturn tail
   ```
   func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    tail := head.Next
    head.Next = swapPairs(head.Next.Next)
    tail.Next = head
    return tail
}
   ```
4. K个一组反转链表
   思路： Dummy node
   prev = dummy k=3
   d -> [1->2->3] -> 4 -> 5

   tail - 1 cur - 2 next -3

   fist iteration (2->1->3)
   cur.Next = prev.Next (2->1)
   prev.Next = cur (d->2)
   tail.Next = next (1->3)
   cur = next
   ```
   func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    dummy := &ListNode{
        0,
        head,
    }
    prev := dummy
    for prev != nil {
        prev = reverse(prev, k)
    }  
    return dummy.Next
}

func reverse(prev *ListNode, k int) *ListNode{
    last := prev
    for i := 0; i < k + 1; i++ {
        last = last.Next
        if i != k && last == nil {
            return nil
        }
    }
    tail := prev.Next
    cur := prev.Next.Next
    for (cur != last) {
        next := cur.Next
        cur.Next = prev.Next
        prev.Next = cur
        tail.Next = next
        cur = next
    }
    return tail
}
   ```

5. 环形链表（判断是否有环） 
    思路：用到快慢指针，快慢指针首先指向头结点
    ```
    slow := head
    fast := head
    for fast!= nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
    ```

6. 环形链表 （判断环的入口）
    思路：同4，首先判断有没有环，如果有环，输出第一次meet的Node
    然后找到环中结点的个数loopCnt
    ```
    node := meetingNode(head)
    if node == nil {
        return
    }
    loopCnt := 1
    pNode1 := node
    for pNode1.Next != node {
        pNode1 = pNode1.Next
        loopCnt++
    }
    ```
    接着用两个指针p1、p2指向head节点，p1先向前进loopCnt步，接着p1和p2同时向前遍历，当第二个指针指向环的入口结点的时候，第一个指针已经围绕着环走了一圈，也回到了环的入口结点，因为两个结点相遇，这个节点就是入环的结点

```
func detectCycle(head *ListNode) *ListNode {
    node := meetingNode(head)
    if node == nil {
        return nil
    }
    loopCount := 1
    pNode := node
    for pNode.Next != node {
        loopCount++
        pNode = pNode.Next
    }
    pNode1 := head
    for i:= 0; i < loopCount; i++ {
        pNode1 = pNode1.Next
    }
    pNode2 := head
    for pNode1 != pNode2 {
        pNode1 = pNode1.Next
        pNode2 = pNode2.Next
    }
    return pNode1

}

func meetingNode(head *ListNode) *ListNode {
    if head == nil || head.Next ==  nil {
        return nil
    }
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return slow
        }
    }
    return nil
}
```



7. 删除排序链表中的重复元素
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
8. 删除排序链表中的重复元素 II
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
9. 删除链表的节点
https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
10. 删除链表的第K个节点
https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
11. 反转链表
https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
12 剑指 Offer 35. 复杂链表的复制
https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
13. 链表的中间结点
https://leetcode-cn.com/problems/middle-of-the-linked-list/
14. 反转链表 II
https://leetcode-cn.com/problems/reverse-linked-list-ii/
15.删除链表的倒数第N个节点
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
16. 剑指 Offer 35. 复杂链表的复制
https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
17. 剑指 Offer 52. 两个链表的第一个公共节点
https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

*递归 && DFS && BFS && 分治*

