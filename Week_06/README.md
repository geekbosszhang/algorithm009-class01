动态规划学习笔记

1. define subproblem 找重复性子问题
2. guess part of the solutions 定义状态空间(状态数组)
3. relate subproblem solutions 
4. Recursion & memorize Or build DP table bottom-up. 定义DP状态方程

练习题
1. 不同路径

思路1：创建一个二维数组，先将第一排第一列的item赋值为1，
然后使用递推公式dp[i][j] = dp[i-1][j] + dp[i][j-1]
最后return dp[m-1][n-1]
思路2：二维变为一维
```
func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if i ==  0 || j == 0 {
                dp[i][j] = 1
            } else {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }
    return dp[m-1][n-1]
}
```
2. 不同路径 II （有障碍物）

```
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    width := len(obstacleGrid[0])
    dp := make([]int, width)
    dp[0] = 1
    for _, row := range obstacleGrid {
        for j := 0; j < width; j++ {
            if row[j] == 1 {
                dp[j] = 0
            } else if j > 0 {
                dp[j] += dp[j-1]
            }
        }
    }
    return dp[width-1]
}
```
3. 最长公共子序列

思路一：二维数组

```
func longestCommonSubsequence(text1 string, text2 string) int {
    text1Len := len(text1)
    text2Len := len(text2)
    if text1Len < 1 || text2Len < 1 {
        return 0
    }
    dp := make([][]int, text1Len+1)
    for i, _ := range dp {
        dp[i] = make([]int, text2Len+1)
    }

    for i, v1 := range text1 {
        for j, v2 := range text2 {
            if v1 == v2 {
                dp[i+1][j+1] = dp[i][j] + 1
            } else {
                dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }

    return dp[text1Len][text2Len]
}
```

