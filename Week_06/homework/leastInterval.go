package homework

import "sort"

func leastInterval(tasks []byte, n int) int {
    if len(tasks) < 1 {
        return 0
    }
    dp := make([]int, 26)
    for _, v := range tasks {
        dp[v-'A']++
    }
    count := 0
    total := len(tasks)
    for total > 0 {
        sort.Ints(dp)
        for i := 0; i <= n; i++ {
            if i < 26 && dp[25-i] != 0 {
                dp[25-i]--
                total--
            }
            if total > 0 {
                count++
            }
        }
    }
    return count+1
}