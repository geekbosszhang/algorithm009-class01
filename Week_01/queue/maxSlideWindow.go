package queue

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 {
        return nil
    }
    var deque = make([]int, 0, len(nums))
    res := make([]int, len(nums)-k+1)
    for i := 0; i < len(nums); i++ {
        for len(deque) != 0 && nums[i] >= nums[deque[len(deque) - 1]] {
            deque = deque[:len(deque) - 1]
        }
        deque = append(deque, i)
        if deque[0] == i - k {
            deque = deque[1:]
        }

        if i+1-k >= 0 {
            res[i+1-k]=nums[deque[0]]
        }
    }
    return res
}