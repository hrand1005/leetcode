/*
func productExceptSelf(nums []int) []int {
    n := len(nums)
    
    forward := make([]int, n)
    forward[0] = nums[0]
    for i := 1; i < n; i++ {
        forward[i] = forward[i-1] * nums[i]
    }
    
    backward := make([]int, n)
    backward[n-1] = nums[n-1]
    for i := n-2; i >= 0; i-- {
        backward[i] = backward[i+1] * nums[i]
    }
    
    result := make([]int, n)
    result[0] = backward[1]
    result[n-1] = forward[n-2]
    for i := 1; i < n-1; i++ {
        result[i] = forward[i-1] * backward[i+1]
    }
    
    return result
}
*/

func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    
    product := 1
    for i := 0; i < n; i++ {
        result[i] = product * nums[i]
        product *= nums[i]
    }
    
    product = 1
    for i := n-1; i > 0; i-- {
        result[i] = product * result[i-1]
        product *= nums[i]
    }
    result[0] = product
    
    return result
}