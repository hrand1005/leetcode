func containsNearbyDuplicate(nums []int, k int) bool {
    seen := make(map[int]int)
    for i, n := range nums {
        if j, ok := seen[n]; ok && abs(i-j) <= k {
            return true
        }
        seen[n] = i
    }
    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}