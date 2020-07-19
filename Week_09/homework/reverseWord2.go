package homework

import "strings"

// TODO use two points to solve this problem
func reverseWords2(s string) string {
    wordArr := strings.Fields(s)
    var result string
    for i, v := range wordArr {
        s := reverseFunc(v)
        if i >= 1 {
            result += " "
        }
        result += s
        
    }
    return result
}

func reverseFunc(str string) string {
    s := []rune(str)
    for i:= 0; i < len(s) / 2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
    return string(s)
}