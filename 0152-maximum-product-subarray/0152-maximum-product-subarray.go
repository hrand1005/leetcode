/*
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
*/

func maxProduct(nums []int) int {
    localMin := nums[0]
    localMax := nums[0]
    maxProduct := nums[0]
    for i := 1; i < len(nums); i++ {
        prevMin := localMin
        prevMax := localMax
        localMin = min(prevMin*nums[i], nums[i], prevMax*nums[i])
        localMax = max(prevMin*nums[i], nums[i], prevMax*nums[i])
        maxProduct = max(maxProduct, localMax)
    }
    return maxProduct
}

func min(nums ...int) int {
    min := nums[0]
    for _, n := range nums {
        if n < min {
            min = n
        }
    }
    return min
}

func max(nums ...int) int {
    max := nums[0]
    for _, n := range nums {
        if max < n {
            max = n
        }
    }
    return max
}