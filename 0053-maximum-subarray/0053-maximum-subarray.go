/*
func maxSubArray(nums []int) int {
    maxSum, curSum := nums[0], nums[0]
    
    for _, n := range nums[1:] {
        curSum = max(curSum + n, n) 
        maxSum = max(maxSum, curSum)
    }
    
    return maxSum
}
*/

func maxSubArray(nums []int) int {
    sumTable := make([]int, len(nums))
    sumTable[0] = nums[0]
    
    for i := 1; i < len(nums); i++ {
        sumTable[i] = max(sumTable[i-1] + nums[i], nums[i])
    }
    
    maxSum := nums[0]
    for _, n := range sumTable {
        maxSum = max(maxSum, n)
    }
    
    return maxSum
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}