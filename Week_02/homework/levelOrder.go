package homework

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