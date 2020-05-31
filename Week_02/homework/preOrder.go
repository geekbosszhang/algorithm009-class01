package homework

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