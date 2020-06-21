package hash

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    hashMap := map[string]int{}
    for i:=0; i < len(s); i++ {
        hashMap[string(s[i])]++
        hashMap[string(t[i])]--
    }
    for _, v := range hashMap {
        if v != 0 {
            return false
        }
    }
    return true
}