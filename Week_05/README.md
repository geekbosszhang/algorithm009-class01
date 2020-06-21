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

   思路1：
      head -> 1 -> 2 -> 3 -> 4
      p - head, cur - 1,
   * step 1. p.next = cur.next  (head -> 2)
   * step 2. cur.next = cur.next.next (1 -> 3)
   * step 3. p.next.next = cur (2 -> 1)
   * step 4. p = cur
   * step 5. cur = cur.next
   
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
   * // 1 -> 2 -> 3 -> 4
   * // 2 -> 1 -> 4 -> 3
   * // (2 tail) -> (1 head) -> swapPairs(3 -> 4)
   * // rturn tail

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
   
   ```
   cur.Next = prev.Next (2->1)
   prev.Next = cur (d->2)
   tail.Next = next (1->3)
   cur = next
   ```
   
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

1. 爬楼梯问题
   
   动态规划
 ```
   dp := make([]int, n+1)
   dp[1] = 1
   dp[2] = 2
   for i := 3; i <= n ; i++ {
      dp[i] = dp[i-1] + dp[i-2]
   }
   return dp[n]
```
2. Invert Binary Tree
  思路： 如果当前节点为空，则终止， 建立一个temp, 交换左右节点，左右节点各自继续作为新的根节点进行子节点的左右交换
  ```
   func invertTree(root *TreeNode) *TreeNode {
     if root == nil {
         return nil
     }
     root.Left, root.Right = root.Right, root.Left
     invertTree(root.Left)
     invertTree(root.Right)
     return root
 }
  ```
  
3. validate binary search Tree
  
  思路一：递归检查当前元素是否在上下界中
   ```
   func isValidBST(root *TreeNode) bool {
	   return helper(root, math.MinInt64, math.MaxInt64)
   }

   func helper(root *TreeNode, low int, high int) bool {
      if root == nil {
         return true
      }
      if root.Val <= low || root.Val >= high {
         return false
      }
      return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
   }
   
   ```

  思路2：中序遍历，比较上一个元素与当前元素，二叉搜索树的中序遍历应该会有一个有序的
  
  ```
  func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
	inOrder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		if root.Val <= inOrder {
			return false
		}
		inOrder = root.Val
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return true
}
  ```
  
 3. Binary Tree Preorder Traversal
 
 思路一: 用递归
 ```
 func preorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    preorder(root, &res)
    return res
}

func preorder(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    *res = append(*res, root.Val)
    preorder(root.Left, res)
    preorder(root.Right, res)
}
```
 
 思路二: 用栈实现
 
 ```
 func preorderTraversal(root *TreeNode) []int {
	 if root == nil {
		 return nil
	 }
	 stack := []*TreeNode{}
	 res := []int{}
	 stack = append(stack, root)
	 for len(stack) > 0 {
		 p := stack[len(stack) - 1]
		 res = append(res, p.Val)
		 stack = stack[:len(stack)-1]
		 if p.Right != nil {
			 stack = append(stack, p.Right)
		 }
		 if p.Left != nil {
			 stack = append(stack, p.Left)
		 }
	 }
	 return res
 }
 ```
 
 4. Binary Tree Inorder Traversal
 
 思路一： 用递归
 ```
  func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    inorder(root, &res)
    return res
 }

func inorder(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, res)
    *res = append(*res, root.Val)
    inorder(root.Right, res)
}
```
 
 思路二： 用栈
 
 ```
 func inorderTraversal(root *TreeNode) (res []int) {
    stack := []*TreeNode{}
    cur := root
    for len(stack) > 0 || cur != nil {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        // pop out one element
        cur = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, cur.Val)
        // turn the direction
        cur = cur.Right
    }
    return
}
```
 
 5. Binary Tree PostOrdeer Traveral
 
 思路一： 用递归
 
 ```
 func postorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    postorder(root, &res)
    return res
}

func postorder(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    postorder(root.Left, res)
    postorder(root.Right, res)
    *res = append(*res, root.Val)
}

```
 
 思路二： 用递归实现，困难点在于根节点需要优先于左右孩子进栈，因此需要借助一个reverse的栈，根->右->左
 
 ```
 func postorderTraversal(root *TreeNode) (res []int) {
    if root == nil {
        return
    }
    stack := []*TreeNode{}
    stackForReverse := []int{}
    stack = append(stack, root)
    for len(stack) > 0 {
        top := stack[len(stack)-1]
        stackForReverse = append(stackForReverse, top.Val)
        stack = stack[:len(stack)-1]
        if top.Left != nil {
            stack = append(stack, top.Left)
        }
        if top.Right != nil {
            stack = append(stack, top.Right)
        }
    }
    //Reverse the stackForReverse, which is parent > right > left
    for i := len(stackForReverse)-1; i >= 0; i-- {
        res = append(res, stackForReverse[i])
    }
    return res
}
```
 6. Binary Tree level Traversal
 
 思路： 逐层处理当前节点，借助辅助queue来方便操作先来后到的子层级
 ```
 func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    var res [][]int
    queue := []*TreeNode{root}
    for i := 0; len(queue) > 0; i++ {
        res = append(res, []int{})
        p := []*TreeNode{}
        for j := 0; j < len(queue); j++ {
            node := queue[j]
            res[i] = append(res[i], node.Val)
            if node.Left != nil {
                p = append(p, node.Left)
            }
            if node.Right != nil {
                p = append(p, node.Right)
            }
        }
        queue = p
    }
    return res
}
```
 7. N-ary Tree Level Order Traversal
 
 思路： 同6， 借助辅助queue
 ```
 func levelOrder(root *Node) (result [][]int) {
    if root == nil {
        return
    }
    queue := []*Node{}
    queue = append(queue, root)
    var level int
    for len(queue) > 0 {
        counter := len(queue)
        result = append(result, []int{})
        for i := 0; i < counter; i++ {
            if queue[i] != nil {
                result[level] = append(result[level], queue[i].Val)
                for _, v := range queue[i].Children {
                    queue = append(queue, v)
                }
            }      
        }
        queue = queue[counter:]
        level++
    }
    return result
}
 ```
 8. N-ary Tree Postorder Traversal
 
 思路一： 递归
 ```
 func postorder(root *Node) []int {
    res := make([]int, 0)
    post(root, &res)
    return res
}

func post(root *Node, res *[]int) {
    if root == nil {
        return
    }
    for _, v := range root.Children {
        post(v, res)
    }
    *res = append(*res, root.Val)
}
 ```
 思路二： 迭代法，借助栈和reverse stack Need to rethink the steps
 ```
 func postorder(root *Node) (res []int) {
    if root == nil {
        return nil
    }
    stack := []*Node{}
    stackReverse := []int{}
    stack = append(stack, root)
    for len(stack) > 0 {
        p := stack[len(stack)-1]
        if p.Children != nil {
            stack = stack[:len(stack)-1]
        }
        stackReverse = append(stackReverse, p.Val)
        for i:= 0; i < len(p.Children);  i++ {
            stack = append(stack, p.Children[i])
        }
    }
    for i:= len(stackReverse)-1; i>=0; i-- {
        res = append(res, stackReverse[i])
    }
    return res
}
 ```
 9. N-ary Tree Preorder Traversal
 
 思路一： 迭代法
 ```
 func preorder(root *Node) []int {
    if root == nil {
		return nil
	}
	stack := []*Node{}
	res := []int{}
	stack = append(stack, root)
	for len(stack) > 0 {
		p := stack[len(stack) - 1]
		res = append(res, p.Val)
		stack = stack[:len(stack) - 1]
		for i:= len(p.Children) - 1; i >= 0; i-- {
			stack = append(stack, p.Children[i])
		}
	}
	return res
}
```

思路二：递归
```
func preorder(root *Node) []int {
    res := make([]int, 0)
    preTraverse(root, &res)
    return res
}

func preTraverse(root *Node, res *[]int) {
    if root == nil {
        return
    }
    *res = append(*res, root.Val)
    for _, v :=range root.Children {
        preTraverse(v, res)
    }
}
```
 10. Maximum Depth of Binary Tree
 
 思路： max(left, right) + 1
 
 ```
 func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	depth := 1
	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	if leftHeight > rightHeight {
		depth = leftHeight + 1
	} else {
        depth = rightHeight + 1
    }
	return depth
}
 ```
 11. Minimun Depth of Binary Tree
 
 思路： 比较左右子树谁小，要考虑如果树没有或者只有一个节点
 if left == 0 || right == 0 return left + right + 1
 
 ```
 func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    depth := math.MaxInt64
    if root.Left != nil {
        depth = min(minDepth(root.Left), depth)
    }
    if root.Right != nil {
        depth = min(minDepth(root.Right), depth)
    }
    return depth+1
}
 ```
 12. Lowest Common Ancestor
 
 思路一： 如果当前节点已找到或者为nil, 返回给上一层
 查看左边的节点
 查看右边的节点
 如果左右两边都找到了节点，返回
 说明当前节点就是他们的公共祖先，否则返回左或者右节点（谁不为Null返回谁）
 ```
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     if root == nil {
         return nil
     }
     if root.Val == p.Val || root.Val == q.Val {
         return root
     }
     // terminater if root subtree contains p and q, and p and q not in the same subtree, return root
     left := lowestCommonAncestor(root.Left, p, q)
     right := lowestCommonAncestor(root.Right, p, q)
     if left != nil && right != nil {
         return root
     }
     if left == nil {
         return right
     }
    return left
}
 ```
 13. Construct Binary Tree from Preorder and Inorder Traversal
 
 思路： preorder 找根节点，inorder找到跟及诶点的位置，根节点左边为Left，右边为right
 
 ```
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) < 1 {
        return nil
    }
    root := &TreeNode {
        Val: preorder[0],
        Left: nil,
        Right: nil,
    }
    var rootIndex int
    // Find the root index in inorder and seperate left tree and right tree
    for i, v := range inorder {
        if v == root.Val {
            rootIndex = i
            break
        }
    }
    root.Left = buildTree(preorder[1:len(inorder[:rootIndex])+1], inorder[:rootIndex])
    root.Right = buildTree(preorder[len(inorder[:rootIndex])+1:], inorder[rootIndex+1:])
    return root
}
 ```
 14. Find Largest Value in each Tree row
 
 思路 参考题6， 二叉树的按层遍历
 ```
 func largestValues(root *TreeNode) (res []int) {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    for i := 0 ; len(queue) > 0; i++ {
        max := math.MinInt32
        q := []*TreeNode{}
        for j := 0; j < len(queue); j++ {
            node := queue[j]
            if node.Val > max {
                max = node.Val
            }
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        queue = q
        res = append(res, max)
    }
    return res
}
```
  
 15. Generate parentheses
 
 思路是左括号没有达到限制就添加左括号，左括号数量大于右括号就添加右括号，左右两边括号数量都达到限制代表终止i
 ```
 func generateParenthesis(n int) []string {
    result = []string{}
    _generateParenthesis(0, 0, n, "")
    return result
}

func _generateParenthesis(left int, right int, n int, s string) {
    if left == n && right == n {
        result = append(result, s)
    }
    if left < n {
        _generateParenthesis(left + 1, right, n, s + "(")
    }
    if left > right {
        _generateParenthesis(left, right + 1, n, s + ")")
    }
}
 ```
 16. Number of islands
 
 Need rethink
 ```
 func numIslands(grid [][]byte) int {
    nr := len(grid)
    if nr < 1 {
        return 0
    }
    nc := len(grid[0])

    var count = 0
    for i := 0; i < nr; i++ {
        for j := 0; j < nc; j++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j)
                count++
            }
        }
    }
    return count
}

func dfs(grid [][]byte, i int, j int) {
    for i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
        return
    }
    grid[i][j] = '0'
    dfs(grid, i+1, j)
    dfs(grid, i - 1, j)
    dfs(grid, i, j+1)
    dfs(grid, i, j -1)
}
```
 
 17. Minimun Genetic Mutation
 
 ```
 func minMutation(start string, end string, bank []string) int {
    if indexOf(end, bank) == -1 {
        return -1
    }
    queue := []string{start}
    isUsed := make([]bool, len(bank))
    count := 0

    for len(queue) > 0 {
        l := len(queue)
        for i := 0; i < l; i++ {
            cur := queue[i]
            if cur == end {
                return count
            }
            for j := 0; j < len(cur); j++ {
                for _, s := range mutationMap[cur[j]] {
                    if idx := indexOf(cur[:j]+s+cur[j+1:], bank); idx != -1 && !isUsed[idx] {
                        queue = append(queue, bank[idx])
                        isUsed[idx]=true
                    }
                }
            }
        }
        count++
        queue = queue[1:]
    }
    return -1
}
 ```
 
 18. word ladder
 
 思路的核心是BFS
 使用queue，讲start word入队，循环check queue是否为空，队列不为空则出队，对出队的单词进行每个字母从a-z的字符转换，将新单词在列表中进行比较，如果找到end word，则 返回level+1, 如果在词典中，将新单词放入queue中 并将词典单词移除，避免重复处理
 
 19. Pow(x, n)
 
 思路是 不断分一半，终止条件为n二分为0， 同时要考虑n为负数 
 ```
 func myPow(x float64, n int) float64 {
    if n >= 0 {
        return quickPow(x, n)
    }
    return 1.0 / quickPow(x, -n)

}

func quickPow(x float64, n int) float64{
    if n == 0 {
        return 1
    }
    y := quickPow(x, n / 2)
    if n % 2 == 0 {
        return y * y
    }
    return y * y * x
}
 ```
 20. Majority Element
 
 思路： 可以通过hashMap或者一遍排序，然后输出 nums[len/2]
 
 21. letter combinations of a phone number
 
 ```
 func letterCombinations(digits string) []string {
    numberMap := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    res := []string{}
    if digits == "" {
        return res
    }
    backTrace("", digits, 0, &res, numberMap)
    return res
}

func backTrace(s string, digits string, i int, res *[]string, numberMap map[byte]string) {
    value := numberMap[digits[i]]
	if i == len(digits)-1 {
		for _, v := range value {
			*res = append(*res, s + string(v))
		}
		return
	}
	for _, v := range value {
		backTrace(s + string(v), digits, i+1, res, numberMap)
	}

}
 ```
 22. Permutations
 
 思路： 遍历数组的所有元素，不断drill down到nums.length层，每一层的结果 遇到已在state的元素需要skip, used数组来判重
 以下解法可以优化，通过传地址的方式，也可以track不用清空
 ```
 var res [][]int
func permute(nums []int) [][]int {
    res = [][]int{}
    track := []int{}
    used := make([]bool, len(nums))
    backstrace(nums, track, used)
    return res
}

func backstrace(nums []int, track []int, used []bool) {
    if len(track) == len(nums) {
        //排列结束, 此处一定要copy slice之后再放入结果集！！!
        tmp := make([]int, len(nums))
        copy(tmp, track)
        res = append(res, tmp)
        return
    }
    for i:=0 ; i< len(nums); i++ {
        if !used[i]{
            used[i] = true
            track = append(track,nums[i])
            backstrace(nums, track, used)
            track = track[:len(track)-1]
            used[i] = false
        }

    }
}
 ```
 23. Combinations
 
 思路：DFS从1到N，排列组合优先DFS
 
 ```
 func combine(n int, k int) [][]int {
    ans := [][]int{}
    arr := []int{}
    dfs(1, n, k, arr, &ans)
    return ans
}

func dfs(s, n, k int, cur []int, ans *[][]int) {
    if k < 0 { return }
    if k == 0 { *ans = append(*ans, append([]int{}, cur...)); return }
    for i := s; i <= n; i++ {
        dfs(i + 1, n, k - 1, append(cur, i), ans)
    }
}
 ```
 24. subset
 
 对当前层可选可不选，然后进入下一层
 ```
 func subsets(nums []int) [][]int {
   res := make([][]int, 0)
   solve(0, nums, nil, &res)
   return res
}

func solve(index int, nums []int, cur []int, res *[][]int) {
    if index == len(nums) {
        *res = append(*res, append([]int{}, cur...))
        return
    }
    solve(index+1, nums, cur,  res)
    solve(index+1, nums, append(cur, nums[index]), res)
}
 ```
 25. N queens
 
 TODO 
 
 *Hash*
 
 1. Two sum
 2. 异位词
 
 *Stack*
 1. valid parentheses
 2. min stack
 3. trapping rain water
 4. largest rectangles in histogram
 
 *Queue*
 1. Sliding window Maximum
 2. 循环队列
 
 *Array*
 1. Two sum
 2. Three sum
 3. Remove duplicates from sorted array
 4. Merge sorted array
 5. 移动0到数组最后
 6. Intersection of two arrays
 7. 装水最多的水桶
 8. Plus one
 9. Rotate Array
 10. Matrix
 
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  

