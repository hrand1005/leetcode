func canJump(nums []int) bool {
    maxReach := 0
    
    for i, v := range nums {
        if maxReach < i {
            return false
        }
        maxReach = max(maxReach, i+v)
    }
    
    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}