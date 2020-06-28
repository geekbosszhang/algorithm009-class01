package homework

func minPathSum(grid [][]int) int {
    m := len(grid)
    if m < 1 {
        return 0
    }
    n := len(grid[0])
    dp := make([][]int, m)
    
    for i:= 0; i < m; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                dp[i][j] = grid[0][0]
            }
            if i == 0 && j >= 1{
                dp[0][j] = dp[0][j-1] + grid[i][j]
            }
            if j == 0 && i >= 1 {
                dp[i][0] = dp[i-1][0] + grid[i][j]
            }
            if i >= 1 && j >= 1 {
                dp[i][j]=min(dp[i-1][j], dp[i][j-1])+grid[i][j]
            }
            
        }
    }
    return dp[m-1][n-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}