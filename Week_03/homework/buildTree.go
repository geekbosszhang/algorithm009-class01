package homework

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