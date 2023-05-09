func longestConsecutive(nums []int) int {
    seen := make(map[int]bool)
    for _, n := range nums {
        seen[n] = true
    }
    longest := 0
    for _, n := range nums {
        if !seen[n-1] {
            seq := 1
            for seen[n+seq] {
                seq++
            }
            longest = max(longest, seq)
        }
    }
    return longest
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}