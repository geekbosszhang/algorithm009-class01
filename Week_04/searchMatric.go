package Week_04

func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    low := 0
    high := m * n - 1
    for low <= high {
        mid := low + (high - low) / 2
        element := matrix[mid/n][mid%n]
        if element == target {
            return true
        }
        if element < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
       
    return false
}