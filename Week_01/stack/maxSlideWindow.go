package stack

// 维护一个单调递减的队列,队列中保存对应元素的index 
// TODO not too clear
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	var Q = make([]int, 0, len(nums))
	res := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		for len(Q) != 0 && nums[i] >= nums[Q[len(Q)-1]] {
			Q = Q[:len(Q)-1]
		}
		// 当前元素入栈
		Q = append(Q, i)

		// 窗口离开first元素时，栈中第一个元素出栈
		if Q[0] == i-k {
			Q = Q[1:]
		}
		// 窗口中满了k个数
		if i+1-k >= 0 {
			res[i+1-k] = nums[Q[0]]
		}
	}
	return res
}