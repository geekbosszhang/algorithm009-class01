package homework

func solveSudoku(board [][]byte)  {
    if len(board) < 1 {
        return
    }
    solvedfs(board)
}

func solvedfs(board [][]byte) bool{
	var c byte
	// why using len(board) and 9 both works?
    for i:=0; i < len(board); i++ {
        for j:=0; j < len(board[0]); j++ {
            if board[i][j] == '.' {
                for c = '1'; c <= '9'; c++ {
                    if isValid(board, i, j, c) {
                        board[i][j] = c
                        if solvedfs(board) {
                            return true
                        } else {
                            board[i][j] = '.'
                        }
                    }
                }
                return false
            }
        }
    }
    return true
}

func isValid(board [][]byte, row, col int, c byte) bool{
    for i:= 0; i < 9; i++ {
        if board[i][col] != '.' && board[i][col] == c {
            return false
        }
        if board[row][i] != '.' && board[row][i] == c {
            return false
        }
        if board[3 * (row / 3) + i / 3][3 * (col / 3) + i % 3] != '.' && board[3 * (row / 3) + i / 3][3 * (col / 3) + i % 3] == c {
            return false
        }
    }
    return true
}