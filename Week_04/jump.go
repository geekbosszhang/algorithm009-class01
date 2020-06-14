package Week_04

func jump(nums []int) int {
    step := 0
    maxPostion := 0
    end := 0
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] + i > maxPostion {
            maxPostion = nums[i] + i
        }
        if i == end {
            end = maxPostion
            step++
        }
    }
    return step
}