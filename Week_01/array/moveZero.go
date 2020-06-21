package array

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
        return
    }
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[j] = nums[i]
            j++
        }
    }
    for k := j; k < len(nums); k++ {
        nums[k] = 0
    }
    return
}