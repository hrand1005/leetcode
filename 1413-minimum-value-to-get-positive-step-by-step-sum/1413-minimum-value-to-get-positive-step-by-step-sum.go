func minStartValue(nums []int) int {
    minVal := nums[0]
    total := nums[0]
    for _, n := range nums[1:] {
        total += n
        minVal = min(minVal, total)
    }
    return 1 - min(minVal, 0)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}