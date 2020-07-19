package homework

func reverseOnlyLetters(S string) string {
    if len(S) < 1 {
        return ""
    }
    str := []rune(S)
    l := 0
    r := len(str) - 1
    for l < r {
        if !isWord(str[l]) {
            l++
            continue
        }
        if !isWord(str[r]) {
            r--
            continue
        }
        str[l], str[r] = str[r], str[l]
        l++
        r--
    }
    return string(str)
}

func isWord(str rune) bool {
    return str >= 'a' && str <= 'z' || str >= 'A' && str <= 'Z'
}