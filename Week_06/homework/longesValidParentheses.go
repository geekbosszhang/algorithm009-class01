package homework

func longestValidParentheses(s string) int {
    stack := make([]rune, 0)
    dp := make([]int, len(s)+1)
    dp[0] = 0
    max := dp[0]
    for i, v := range s {
        index := i + 1
        if v == '(' {
            stack = append(stack, v)
        } else {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
                pairs := 1 + dp[index-1]
                prev_index := index - pairs * 2
                if prev_index > 0 {
                    pairs += dp[prev_index]
                }
                dp[index] = pairs
                if dp[index] > max {
                    max = dp[index]
                }
            }
        }
    }
    return max * 2
}