package homework

func findCircleNum(M [][]int) int {
    if len(M) < 1 {
        return 0
    }
    n := len(M)
    p := make([]int, n)
    for i:= 0; i < n; i++ {
        p[i] = i
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if M[i][j] == 1 {
                union(p, i, j)
            }
        }
    }
    count := 0
    hashMap := map[int]bool{}
    for i, _ := range p {
        parentValue := parent(p, i)
        if _, ok := hashMap[parentValue]; !ok{
            count++
            hashMap[parentValue] = true
        }
    }
    return count
}

func union(p []int, i, j int) {
    p1 := parent(p, i)
    p2 := parent(p, j)
    p[p2] = p1
}

func parent(p []int, i int) int{
    root := i
    for p[root] != root {
        root = p[root]
    }
    // 路径压缩
    for p[i] != i {
        i, p[i] = p[i], root
    }
    return root
}