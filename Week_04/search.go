package Week_04

func search(nums []int, target int) int {
    if len(nums) < 1 {
        return -1
    }
    start := 0
    end := len(nums) - 1
    mid := 0
    for start <= end {
        mid = start + (end-start)/2
        if nums[mid] == target {
            return mid
        }
        if nums[start] <= nums[mid] {
            if target >= nums[start] && target < nums[mid] {
                end = mid -1
            } else {
                start = mid + 1
            }
        } else {
            if target <= nums[end] && target > nums[mid] {
                start = mid + 1
            } else {
                end = mid -1
            }
        }
    }
    return -1
}