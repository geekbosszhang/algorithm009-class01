package tree

func levelOrder(root *Node) (result [][]int) {
    if root == nil {
        return result
    }
    var level int
    queue := []*Node{}
    queue = append(queue, root)
    for len(queue) > 0 {
		counter := len(queue)
		// Need to clean up the result
        result = append(result, []int{})
        for i:= 0; i < counter; i++ {
			// Need to check if queue[i] is nil
            if queue[i] != nil {
                result[level] = append(result[level], queue[i].Val)

                for _, v := range queue[i].Children{
                    queue = append(queue, v)
                }
            }
        }
        queue = queue[counter:]
        level++   
    }
    return result

}