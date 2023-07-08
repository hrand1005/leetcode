func longestSubarray(nums []int) int {
    r, l, zeroes := 0, 0, 0
    for ; r < len(nums); r++ {
        if nums[r] == 0 {
            zeroes++
        } 
        if zeroes > 1 {
            if nums[l] == 0 {
                zeroes--
            }
            l++
        }
    }
    return r - l - 1
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