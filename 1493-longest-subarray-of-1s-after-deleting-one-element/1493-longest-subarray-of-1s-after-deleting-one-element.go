func longestSubarray(nums []int) int {
    right, left, zeroCount := 0, 0, 0
    for right < len(nums) {
        if nums[right] == 0 {
            zeroCount++
        } 
        right++
        
        if zeroCount > 1 {
            if nums[left] == 0 {
                zeroCount--
            }
            left++
        }
    }
    return right - left - 1
}

/*
func longestSubarray(nums []int) int {
    zeroFound := false
    before, after, maxSeq := 0, 0, 0
    for _, n := range nums {
        if n == 0 {
            zeroFound = true
            maxSeq = max(maxSeq, before + after)
            before, after = after, 0
        } else {
            after++
        }
    }
    maxSeq = max(maxSeq, before + after)
    if zeroFound {
        return maxSeq
    }
    return maxSeq-1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
*/