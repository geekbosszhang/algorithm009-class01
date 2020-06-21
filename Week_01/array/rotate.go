package array

func rotate(nums []int, k int)  {
    arrLen := len(nums)
    if arrLen < 1 {
        return 
	}
	// if k > arrLen
    k %= arrLen
    reverse(nums, 0, arrLen - 1)
    reverse(nums, 0, k - 1)
    reverse(nums, k, arrLen - 1)
}

func reverse(nums []int, start int, end int) {
    if len(nums) < 2 {
        return
    }
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}