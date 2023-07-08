func longestOnes(nums []int, k int) int {
    seq := 0
    maxSeq := 0
    
    for i, _ := range nums {
        seq = calculateSeq(nums[i:], k)
        maxSeq = max(maxSeq, seq)
    }
    return maxSeq
}

func calculateSeq(nums []int, k int) int {
    idx := 0
    for k != 0 && idx < len(nums) {
        if nums[idx] == 0 {
            k--
        }
        idx++
    }
    for idx < len(nums) {
        if nums[idx] == 0 {
            break
        }
        idx++
    }
    return idx
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}