学习笔记

股票问题学习总计
1. K  = 1
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
         = max(dp[i-1][1], 0-prices[i])
         = max(dp[i-1][1], -prices[i])
(because dp[i-1][0][0] = 0, k=1, but buy one means k-1)

i == 0
dp[-1][0] = 0
dp[-1][1] = -inifinity

```
func maxProfit(prices []int) int {
    dp_i_0 := 0
    dp_i_1 := math.MinInt32
    for i:= 0; i < len(prices); i++ {
        dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = max(dp_i_1, -prices[i])
    }
    return dp_i_0
}
```

2. k is infinity 
可以多次买卖一只股票，但是不能同时参与多笔交易 ？？？
思路是K是正无穷，那么K和K-1可以认为是一样的
因此dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
               = max(dp[i-1][k][1], dp[i-1][k][0]-prices[i]
K可以忽略
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])

```
func maxProfit(prices []int) int {
    dp_i_0 := 0
    dp_i_1 := math.MinInt32
    for i:= 0; i < len(prices); i++ {
        temp := dp_i_0
        dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
        dp_i_1 = max(dp_i_1, temp - prices[i])
    }
    return dp_i_0
}
```

3. K = inifinity but has cool time, 卖出后要等一天才可以买入
思路是
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])

dp[i-2][0] => dp_pre_0 need to update the value in the for loop

```
func maxProfit(prices []int) int {
    dp_i_0 := 0
    dp_i_1 := math.MinInt32
    dp_pre_0 := 0
    for i:= 0; i < len(prices); i++ {
        temp := dp_i_0
        dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = max(dp_i_1, dp_pre_0 - prices[i])
        dp_pre_0 = temp
    }
    return dp_i_0
}
```
4. K = inifinity but has fee
思路是参考题二
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i]-fee)
```
func maxProfit(prices []int, fee int) int {
    dp_i_0 := 0
    dp_i_1 := math.MinInt32
    for i:= 0; i < len(prices); i++ {
        temp := dp_i_0
        dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = max(dp_i_1, temp - prices[i] - fee)
    }
    return dp_i_0
}
```
5. K = 2
base case：
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity

状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

```
func maxProfit(prices []int) int {
    maxK := 2
    n := len(prices)
    if n < 1 {
        return 0
    }
    // this is initilize three demention array in go
    dp := make([][][]int, n+1)
    for i := range dp {
        dp[i] = make([][]int, maxK+1)
        for j := range dp[i] {
            dp[i][j] = make([]int, 2)
        }
    }

    for i:= 0; i < n; i++ {
        for k := maxK; k >= 1; k--{
            if i == 0 {
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
                continue
            }
            dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
            dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
        }
    }
    return dp[n-1][maxK][0]
}
```

6. K is any
思路是dp开辟内存太多会导致内存越界的问题，因此需要基于题5和题2进行优化
一次交易由买入和卖出构成，至少需要两天。所以说有效的限制 k 应该不超过 n/2，如果超过，就没有约束作用了，相当于 k = +infinity。这种情况是之前解决过的。

```
 maxK := k
    n := len(prices)
    if n < 1 {
        return 0
    }
    if k > n / 2 {
        return maxProfitKInifinite(k, prices)
    }
    dp := make([][][]int, n+1)
    for i := range dp {
        dp[i] = make([][]int, maxK+1)
        for j := range dp[i] {
            dp[i][j] = make([]int, 2)
        }
    }
    for i:= 0; i < n; i++ {
        for k := maxK; k >= 1; k--{
            if i == 0 {
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
                continue
            }
            dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
            dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
        }
    }
    return dp[n-1][maxK][0]
```

7. 不同路径II问题

```
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i:= 0; i < m && obstacleGrid[i][0] == 0; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
        dp[0][j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if (obstacleGrid[i][j] == 0) {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }
    }
    return dp[m - 1][n - 1];
}
```

字符串基础篇：
1. 字符串toLowerCase
2. 字符串转换成整数atoi
3. 最长公共前缀
4. 反转字符串
5. 字母异位词
6. 回文串

高级字符串和DP
1. 编辑距离
动态规划方程
dp[i][j] 表示word1(0, i) 与 word2(0, j)之间的编辑距离
```
if word1[i] = word2[j] {
    edit_dist(i, j) = edit_dist(i-1, j-1)
} else {
    edit_dist(i, j) = min(edit_dist(i-1, j-1), edit_dist(i-1, j), edit_dist(i, j-1)) + 1
}
```
2. 最长子序列问题VS子串
子串是必须要连续的，子序列不要求连续
dp[i][j] 表示word1与word2之间的最长子序列
```
if s1[i-1] = s2[j-1] {
    dp[i][j] = dp[i-1][j-1] + 1
} else {
    dp[i][j] = max(dp[i-1][j], dp[i][j-1])
}
``` 
最长子串问题
```
if s1[i-1] == s2[j-1] {
    dp[i][j] = dp[i-1][j-1] + 1
} else {
    dp[i][j] = 0
}
```
3. 最长回文子串问题
枚举回文串的中心

