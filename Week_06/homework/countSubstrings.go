package homework

func countSubstrings(s string) int {
    // dp[i][j] = true表示是回文，否则为false
    // if s[i] == s[j] if dp[i+1][j-1]是，则dp[i][j]就是
    // if s[i] != s[j] dp[i][j] = false
    if len(s) < 1 {
        return 0
    }
    result := len(s)
    dp := make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
        dp[i][i] = true
    }
    for i := len(s) -1 ; i >= 0 ; i-- {
        for j := i+1; j < len(s); j++ {
            if s[i] == s[j] {
                if j - i == 1 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            } else {
                dp[i][j] = false
            }
            if dp[i][j] {
                result++
            }
        }
    }
    return result
}