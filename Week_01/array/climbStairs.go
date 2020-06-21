package array

// Use Fibonacci solutions.
func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    oneStepBefore := 2
    twoStepBefore := 1
    for i:= 3; i<= n; i++ {           
		oneStepBefore, twoStepBefore = oneStepBefore + twoStepBefore, oneStepBefore
    }
    return oneStepBefore 
}