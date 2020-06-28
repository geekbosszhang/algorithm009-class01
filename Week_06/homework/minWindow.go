package homework

func minWindow(s string, t string) string {
    if len(s) < 1 || len(t) < 1 || len(s) < len(t) {
        return ""
    }
    // 字符出现的次数
    need, window := map[byte]int{}, map[byte]int{}
    for i := range t {
        need[t[i]]++
    }
    left, right := 0, 0
    // valid保存验证通过的字符数量
    valid := 0
    // index, length 表示子串的起始位置和长度
    index := 0
    length := len(s) + 1
    for right < len(s) {
        // 判断s的子串中包含的有效的T中的子串
        if _, ok := need[s[right]]; ok {
            window[s[right]]++
            if window[s[right]] == need[s[right]] {
                valid++
            }
        }
        right++
        for valid == len(need) {
            if right - left < length {
                index = left
                length = right - left
            }
            if _, ok := need[s[left]]; ok {
                window[s[left]]--
                if window[s[left]] < need[s[left]] {
                    valid--
                }
            }
            // 开始左移
            left++
        }
    }
    if length == len(s) + 1 {
        return ""
    }
    return s[index: index + length]
}