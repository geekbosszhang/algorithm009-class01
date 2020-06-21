package stack

func isValid(s string) bool {
    if len(s) % 2 != 0 {
        return false
    }
    stack := []string{}
    for _, v := range s {
        if string(v) == "(" {
            stack = append(stack, ")")
        } else if string(v) == "[" {
            stack = append(stack, "]")
        } else if string(v) == "{" {
            stack = append(stack, "}")
        } else {
            if len(stack) < 1 {
                return false
            }
            top := stack[len(stack)-1]
            if top == string(v) {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }
    return len(stack) == 0
}