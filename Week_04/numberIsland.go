package Week_04

func numIslands(grid [][]byte) int {
    nr := len(grid)
    if nr < 1 {
        return 0
    }
    nc := len(grid[0])

    var count = 0
    for i := 0; i < nr; i++ {
        for j := 0; j < nc; j++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j)
                count++
            }
        }
    }
    return count
}

func dfs(grid [][]byte, i int, j int) {
    for i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
        return
    }
    grid[i][j] = '0'
    dfs(grid, i+1, j)
    dfs(grid, i - 1, j)
    dfs(grid, i, j+1)
    dfs(grid, i, j -1)
}