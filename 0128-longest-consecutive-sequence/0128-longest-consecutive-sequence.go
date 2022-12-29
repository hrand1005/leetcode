func longestConsecutive(nums []int) int {
    numSet := toSet(nums)
    longestSeq := 0
    for n := range numSet {
        if exists := numSet[n-1]; !exists {
            thisSeq := 1
            for i := n+1; numSet[i]; i++ {
                thisSeq++
            }
            longestSeq = max(longestSeq, thisSeq)
        }
    }
    return longestSeq
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