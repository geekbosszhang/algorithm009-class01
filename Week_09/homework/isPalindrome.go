package homework

import "strings"
func isPalindrome(s string) bool {
    str := []rune(s)
    l := 0
    r := len(str) - 1
    for l < r {
        if strings.ToLower(string(str[l])) == strings.ToLower(string(str[r])) {
            l++
            r--
        } else if !isNumberOrLetter(str[l]) {
            l++
        } else if !isNumberOrLetter(str[r]) {
            r--
        } else {
            return false
        }
    }
    return true
}

func isNumberOrLetter(r rune) bool {
    return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9'
}