package homework

func longestCommonPrefix(strs []string) string {
    if len(strs) < 1 {
        return ""
    }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        prefix = lcp(prefix, strs[i])
        if len(prefix) == 0 {
            break
        }
    }
    return prefix
}

func lcp(prefix, compareStr string) string {
    compareLength := 0
    if len(prefix) < len(compareStr) {
        compareLength = len(prefix)
    } else {
        compareLength = len(compareStr)
    }
    index := 0
    for index < compareLength && prefix[index] == compareStr[index] {
        index++
    }
    return prefix[:index]
}