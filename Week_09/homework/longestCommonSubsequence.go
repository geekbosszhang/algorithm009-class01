package homework

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

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}