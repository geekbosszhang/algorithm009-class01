package homework

type Node struct {
	Val int
	Children []*Node
}

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
