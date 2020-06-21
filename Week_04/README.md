使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方
```
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
```
