package stack

func trap(height []int) int {
    left, leftMax, rightMax, ans := 0, 0, 0, 0
    right := len(height) - 1
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            } else {
                ans += leftMax - height[left]
            }
            left++
        } else {
            if height[right] > rightMax {
                rightMax = height[right]
            } else {
                ans += rightMax - height[right]
            }
            right--
        }
    }
    return ans
}