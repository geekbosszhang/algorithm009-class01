package stack

// Option1 暴力法求解
func largestRectangleArea1(heights []int) int {
    if len(heights) < 1 {
        return 0
    }
    if len(heights) == 1 {
        return heights[0]
    }
    maxArea := 0
    for i := 0; i <= len(heights) - 2; i++ {
        minHeight := heights[i]
        for j := i + 1; j <= len(heights) - 1; j++ {
            // find the min Height
            if heights[j] < minHeight {
                minHeight = heights[j]
            }
            area := minHeight * (j - i + 1)
            if area > maxArea {
                maxArea = area
            }
        }
    }
    return maxArea
}
// [2,1,5,6,2,3] => 10 pass
// [2,0,2] 输出 0 expected 2 fail
// [1, 1] => 2 pass
// [0, 9] 输入 0 expected 9 fail

// Option2 暴力法优化解，找到左边界和右边界
func largestRectangleArea2(heights []int) int {
    if len(heights) < 1 {
        return 0
    }
    if len(heights) == 1 {
        return heights[0]
    }
    maxArea := 0
    for i, v := range heights {
        left := i
        right := i
        for j := i-1; j >= 0 ;j-- {
            // Find first element in the left that less than current height
            if heights[j] < v {
                left = j
                break
            }
        }
        for k := i+1; k < len(heights); k++ {
            // Find first element in the right that less than current height
            if heights[k] < v {
                right = k
                break
            }
        }
        area := v * (right - left - 1)
        if area > maxArea {
            maxArea = area
        }
    }
    return maxArea
}

// [2,1,5,6,2,3] => 10 pass
// [2,0,2] 输出 0 expected 2 fail
// [1, 1] 输出0 expected 2 fail
// [0, 9] 输入 0 expected 9 fail


// Option3 using stack 
func largestRectangleArea3(heights []int) int {
    stack := []int{-1}
    maxArea := 0

    for i, v := range heights {
        for stack[len(stack)-1] != -1 && v <= heights[stack[len(stack)-1]] {
            // pop stack
            peek := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            newTop := stack[len(stack)-1]
            area := heights[peek] * (i - 1 - newTop)
            if area > maxArea {
                maxArea = area
            }
            if newTop == -1 {
                break
            }
        }
        stack = append(stack, i)
    }

    for len(stack) > 0 && stack[len(stack)-1] != -1 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        area := heights[top] * (len(heights)-1-stack[len(stack)-1])
        if area > maxArea {
            maxArea = area
        }
    }
    return maxArea
}
// time limit passed 94/96 cases












