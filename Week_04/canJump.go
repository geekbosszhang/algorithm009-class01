package Week_04

func canJump(nums []int) bool {
    if len(nums) < 1 {
        return false
    }
    rightMost := 0
    for i:=0 ; i < len(nums); i++ {
        if i <= rightMost {
            if nums[i] + i > rightMost {
                rightMost = nums[i]+i
            }
            if rightMost >= len(nums) - 1 {
                    return true
            }
        }
    }
    return false
}