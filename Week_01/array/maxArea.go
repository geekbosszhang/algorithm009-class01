package array

func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxArea := 0
    for left < right {
        var h int
        distance := right - left
        if height[left] < height[right] {
            h = height[left]
            left++
        } else {
            h = height[right]
            right--
        }
        area := distance * h
        if area  > maxArea {
            maxArea = area
        }
    }
    return maxArea
}