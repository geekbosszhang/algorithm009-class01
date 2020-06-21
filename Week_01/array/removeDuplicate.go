package array

func removeDuplicates(nums []int) int {
    arrLen := len(nums)
    if arrLen <= 1 {
        return arrLen
    }
    duplicateCount := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            duplicateCount++
        } else {
            nums[i - duplicateCount] = nums[i]
        }
    }
    return arrLen - duplicateCount
}