package homework

import "strings"
func toLowerCase(str string) string {
    var result strings.Builder
    result.Grow(len(str))
    for i:= 0; i < len(str); i++ {
        c  := str[i]
        if c >= 'A' && c <= 'Z' {
            c += 'a' - 'A'
        }
        result.WriteByte(c)

    }
    return result.String()
}