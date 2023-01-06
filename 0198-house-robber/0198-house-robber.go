func rob(nums []int) int {
    cache := make(map[int]int)
    return robRecursive(0, nums, cache)
}


func robRecursive(idx int, nums []int, cache map[int]int) int {
    if len(nums)-1 < idx {
        return 0
    } 
    if len(nums)-1 == idx {
        return nums[idx]
    }
    if v, ok := cache[idx]; ok {
        return v
    }
    
    robCurrent := nums[idx] + robRecursive(idx+2, nums, cache)
    robNext := robRecursive(idx+1, nums, cache)
    maxProfit := max(robCurrent, robNext)
    cache[idx] = maxProfit
    
    return maxProfit
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}