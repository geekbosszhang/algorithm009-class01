package homework

// Use two pointers
func reverseStr2(s string, k int) string {
    str := []rune(s)
    var i int
    for i < len(str) {
        l := i
        r := i+k-1
        if r >= len(str) {
            r = len(str) - 1
        }
        for l < r && l < len(str) {
            str[l], str[r] = str[r], str[l]
            l++
            r--
        }
        i = i + 2 * k
    }
    return string(str)
}