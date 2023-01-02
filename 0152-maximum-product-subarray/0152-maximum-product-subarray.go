func maxProduct(nums []int) int {
    maxProduct := nums[0]
    current := 1
    for i := 0; i < len(nums); i++ {
        current *= nums[i]
        maxProduct = max(maxProduct, current)
        if current == 0 {
            current = 1
        }
    }
    
    current = 1
    for i := len(nums)-1; i > -1; i-- {
        current *= nums[i]
        maxProduct = max(maxProduct, current)
        if current == 0 {
            current = 1
        }
    }
    
    return maxProduct
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}