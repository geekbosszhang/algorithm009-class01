package homework

import "strconv"

func numDecodings(s string) int {
    // 类比爬楼梯问题
    // f(n) = f(n-1) + f(n-2)(if s[n-2] > 10 && s[n-2] < 26)
    if s[0] == '0' {
        return 0
    }
    memo := make([]int, len(s)+1)
    memo[0] = 1
    memo[1] = 1
    for i := 2; i <= len(s); i++ {
        if s[i-1] != '0' {
            memo[i] = memo[i-1]
        } else {
            memo[i] = 0
        }
        if x, err := strconv.Atoi(s[i-2:i]); err == nil {
            if x >= 10 && x <= 26 {
                memo[i] += memo[i-2]
            }
        }
    }
    return memo[len(s)]
}