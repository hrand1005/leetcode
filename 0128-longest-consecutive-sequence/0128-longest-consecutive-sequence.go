func longestConsecutive(nums []int) int {
    numSet := toSet(nums)
    longest := 0
    for n := range numSet {
        if exists := numSet[n-1]; !exists {
            seq := 1
            for numSet[n+seq] {
                seq++
            }
            longest = max(longest, seq)
        }
    }
    return longest
}

func toSet(s []int) map[int]bool {
    set := make(map[int]bool, len(s)) 
    for _, v := range s {
        set[v] = true
    }
    return set
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}