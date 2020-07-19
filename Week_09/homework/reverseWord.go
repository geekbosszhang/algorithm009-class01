package homework

import "strings"
func reverseWords(s string) string {
    str := strings.Trim(s, " ")
    wordArr := strings.Fields(str)
    if len(wordArr) < 1 {
        return ""
    }
    var newWordArr []string
    for i:= 0; i< len(wordArr); i++ {
        newWordArr = append(newWordArr, wordArr[len(wordArr)-1-i])
    }
    return strings.Join(newWordArr, " ")
}