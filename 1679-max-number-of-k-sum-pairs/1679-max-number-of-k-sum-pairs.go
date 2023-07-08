func maxOperations(nums []int, k int) int {
    nCount := make(map[int]int)
    for _, n := range nums {
        nCount[n]++
    }
    
    ops := 0
    counted := make(map[int]bool)
    for n, count := range nCount {
        if !counted[n] {
            if n != (k-n) {
                ops += min(count, nCount[k-n])
            } else {
                ops += count / 2
            }
            counted[n] = true
            counted[k-n] = true
        }
    }
    return ops
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

/*
func maxOperations(nums []int, k int) int {
    ops := 0
    for len(nums) != 0 {
        found := false
        comps := make(map[int]int)
        for i, n := range nums {
            if v, ok := comps[k-n]; ok {
                nums = append(nums[:i], nums[i+1:]...)
                nums = append(nums[:v], nums[v+1:]...)
                found = true
                ops++
                break;
            }
            comps[n] = i
        }
        if !found {
            break
        }
    }
    return ops
}
*/