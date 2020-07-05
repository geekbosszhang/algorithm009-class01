package homework

func findWords(board [][]byte, words []string) (res []string) {
    // 遍历word， 建立Trie树
    trie := new(Trie)
    for _, v := range words {
        trie.Insert(v)
    }
    // Board DFS search
    for i:= 0; i < len(board); i++  {
        for j := 0; j < len(board[0]); j++ {
            existRecursive(board, i, j, trie, &res);
        }
    }
    return res
}

func existRecursive(board [][]byte, row, col int, trie *Trie, res *[]string) {
    if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
        return
    }
    cur := board[row][col]
    if cur == '$' || trie.children[cur - 'a'] == nil {
        return
    }
    trie = trie.children[cur - 'a']
    if trie.data != "" {
        *res = append(*res, trie.data)
        trie.data = ""
    }
    temp := board[row][col]
    board[row][col] = '$'
    existRecursive(board, row-1, col, trie, res)
    existRecursive(board, row + 1, col, trie, res)
    existRecursive(board, row, col-1, trie, res)
    existRecursive(board, row, col+1, trie, res)
    board[row][col] = temp
}