package homework

var count int
func totalNQueens(n int) int {
    count = 0
    solve(0, 0, 0, 0, n)
    return count
}

func solve(row, col, pie, na, n int) {
    if row >= n{
        count++
        return
    }

    pos := ((1<<uint(n)) - 1) & (^(col | pie | na))
    for pos != 0 {
        p := pos&(-pos)
        solve(row+1,col|p,(pie|p)<<1, (na|p)>>1, n)
        pos = pos&(pos - 1)
    }
}