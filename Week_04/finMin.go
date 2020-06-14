package Week_04

func findMin(nums []int) int {
    i := 0
    j := len(nums) - 1
    for i < j {
        if nums[i] > nums[j] {
            i++
        } else {
            j--
        }
    }
    return nums[i]
}