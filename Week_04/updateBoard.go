package Week_04

func updateBoard(board [][]byte, click []int) [][]byte {
    r, c := click[0], click[1]
    if board[r][c] == 'M' {
        board[r][c] = 'X'
        return board
    }
    update(board, r, c)
    return board
}

func update(b [][]byte, r, c int) {
    if r < 0 || c < 0 || r >= len(b) || c >= len(b[0]) || b[r][c] != 'E' {
        return
    }
    mines := mineCount(b, r, c)
    if mines == 0 {
        b[r][c] = 'B'
        dirs := [][]int{
            []int{0, 1},
            []int{0, -1},
            []int{-1, 0},
            []int{1, 0},
            []int{-1, -1},
            []int{-1, 1},
            []int{1, -1},
            []int{1, 1},
        }
        for _, v := range dirs {
            update(b, r+v[0], c+v[1])
        }
    } else {
        b[r][c] = byte(mines) + '0'
    }
}

func mineCount(b [][]byte, r, c int) int {
    var count int
    for i := r-1; i <= r+1; i++ {
        for j := c -1; j <= c+1; j++ {
            if i < 0 || j < 0 || i >= len(b) || j >= len(b[0]) {
                continue
            }
            if (b[i][j] == 'M' || b[i][j] == 'X') {
                count++
            }
        }
    }
    return count
}