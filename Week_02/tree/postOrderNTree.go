package tree

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