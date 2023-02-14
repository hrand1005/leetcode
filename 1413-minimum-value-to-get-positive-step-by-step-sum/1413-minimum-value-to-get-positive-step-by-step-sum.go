func minStartValue(nums []int) int {
    minVal := 0
    total := 0
    for _, n := range nums {
        total += n
        minVal = min(minVal, total)
    }
    return 1 - minVal
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}