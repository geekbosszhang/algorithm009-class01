package homework

import (
	"strings"
	"math"
)

func myAtoi(str string) int {
    s := strings.Trim(str, " ")
    if len(s) < 1 {
        return 0
    }
    sign := 1
    n := 0
    for i, v := range s {
        if i == 0 {
            if v == '-' {
                sign = -1
            } else if v == '+' {
                sign = 1
            } else if v - '0' >= 0 && v - '0' <= 9 {
                n = n * 10 + int(v - '0')
            } else {
                return n
            }
        } else {
            if v - '0' < 0 || v - '0' > 9 {
                break
            }
            if sign == -1 && (n * 10 + sign * int(v - '0') < math.MinInt32) {
                return math.MinInt32
            }
            if sign == 1 && (n * 10 + int(v - '0') > math.MaxInt32) {
                return math.MaxInt32
            }
            n = n * 10 + sign * int(v - '0')
        }
    }
    return n
}