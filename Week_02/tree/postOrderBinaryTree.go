package tree

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