package homework

func nthUglyNumber(n int) int {
    dp := make([]int, n)
    a, b, c := 0, 0, 0
    dp[0] = 1
    for i:= 1; i < n ; i++ {
        dp[i] = min(2 * dp[a], 3 * dp[b], 5 * dp[c])
        if dp[i] == 2 * dp[a] {
            a++
        }
        if dp[i] == 3 * dp[b] {
            b++
        }
        if dp[i] == 5 * dp[c] {
            c++
        }
    }
    return dp[n-1]
}

func min(x, y, z int) int {
    minValue := x
    if y < x {
        minValue = y
    }
    if z < x {
        minValue = z
    }
    return minValue
}