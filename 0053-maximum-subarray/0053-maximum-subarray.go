func maxSubArray(nums []int) int {
    maxSum, curSum := nums[0], nums[0]
    
    for _, n := range nums[1:] {
        curSum = max(curSum + n, n) 
        maxSum = max(maxSum, curSum)
    }
    
    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}