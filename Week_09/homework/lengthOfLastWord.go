package homework

import "strings"

func lengthOfLastWord(s string) int {
    resArr := strings.Fields(s)
    if len(resArr) < 1 {
        return 0
    }
    return len(resArr[len(resArr)-1])
}