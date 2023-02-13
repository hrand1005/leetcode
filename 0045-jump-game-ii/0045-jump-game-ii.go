func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    start, end := 0, nums[0]
    jumps := 1
    for end < len(nums)-1 {
        jumps++
        scope := 0
        for i := start; i < end+1; i++ {
            scope = max(scope, i+nums[i])
        }
        start, end = end, scope
    }
    return jumps
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}