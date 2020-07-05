package homework

func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    op := make([][]int, m + 1)
    for i := 0; i < m + 1; i++ {
        op[i] = make([]int, n + 1)
        op[i][0] = i
    }
    for j := 0; j < n + 1; j++ {
        op[0][j] = j
    }  

    for i:= 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                op[i][j] = op[i-1][j-1]
            } else {
                op[i][j]= 1 + min(op[i][j-1], min(op[i-1][j], op[i-1][j-1]))
            }
        }
    }
    return op[m][n]
}