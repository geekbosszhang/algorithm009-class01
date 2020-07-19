package homework

func firstUniqChar(s string) int {
    if len(s) < 1 {
        return -1
    }
    hashMap := make(map[string]int, 0)
    for _, v := range s {
        hashMap[string(v)]++
    }
    for key, v2 := range s {
        if hashMap[string(v2)] == 1 {
            return key
        }
    }
    return -1

}