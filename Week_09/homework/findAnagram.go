package homework

func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return nil
    }
    var res []int
    psum := 1
    for i:=0; i<len(p); i++ {           // 计算 p 的字符累乘结果
        psum *= int(p[i])
    }

    for i:= 0; i <= len(s) - len(p); i++ {
        compareWord := s[i: i+len(p)]
        if getSum(compareWord) == psum {
            res = append(res, i)
        }
    }
    return res
}

func getSum(s string) int {   // 判断两个字符串是否为字母异位词
    sum := 1
    for i:=0; i<len(s); i++ {
        sum *= int(s[i])
    }
    return sum
}