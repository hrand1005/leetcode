func longestOnes(nums []int, k int) int {
    zeroes := 0
    l, r := 0, 0
    maxSeq := 0
    for r < len(nums) {
        if nums[r] == 0 {
            zeroes++
        }
        
        if zeroes > k {
            if nums[l] == 0 {
                zeroes--
            }
            l++
        }
        
        maxSeq = max(maxSeq, (r - l) + 1)
        r++
    }
    return maxSeq
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/*
func longestOnes(nums []int, k int) int {
    ones := k
    maxSeq := 0
    l, r := 0, 0
    
    for r < len(nums) {
        if nums[r] == 1 {
            r++
        } else if ones > 0 {
            ones--
            r++
        } else {
            maxSeq = max(maxSeq, r - l)
            l++
            r = l
            ones = k
        }
    }
    
    return max(maxSeq, r - l)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
*/

/*
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
*/