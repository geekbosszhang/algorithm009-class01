package homework

func solve(board [][]byte)  {
    if len(board) < 1 {
        return
    }
    rows := len(board)
    cols := len(board[0])
    p := make([]int, rows * cols + 1)
    for i:= 0; i < rows * cols + 1; i++ {
        p[i] = i
    }
    dummpyNode := rows * cols

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == 'O' {
                // 当前节点在边界就和dummy合并
                if i == 0 || i == rows - 1 || j == 0 || j == cols - 1 {
                    union(p, dummpyNode, i * cols + j)
                } else {
                    if board[i-1][j] == 'O' {
                        union(p, i * cols + j, (i-1) * cols + j)
                    }
                    if board[i+1][j] == 'O' {
                        union(p, i * cols + j, (i+1)*cols + j)
                    }
                    if board[i][j-1] == 'O' {
                        union(p, i * cols + j, i * cols + j -1 )
                    }
                    if board[i][j+1] == 'O' {
                        union(p, i * cols + j, i * cols + j + 1)
                    }
                }
            }
        }
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if parent(p, i*cols + j) == parent(p, dummpyNode) {
                board[i][j] = 'O'
            } else {
                board[i][j] = 'X'
            }
        }
    }
}
