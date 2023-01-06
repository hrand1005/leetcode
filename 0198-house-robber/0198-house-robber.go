/*
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
*/

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }
    
    table := make([]int, len(nums))
    table[0] = nums[0]
    table[1] = max(nums[0], nums[1])
    
    for i := 2; i < len(nums); i++ {
        curProfit := nums[i] + table[i-2]
        prevProfit := table[i-1]
        table[i] = max(curProfit, prevProfit)
    }
    
    return table[len(table)-1]
}