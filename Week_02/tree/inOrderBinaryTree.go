package tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

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