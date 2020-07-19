package homework

func numJewelsInStones(J string, S string) int {
    count := 0
    for _, v1 := range J {
        for _, v2 := range S {
            if v2 == v1 {
                count++
            }
        } 
    }
    return count
}