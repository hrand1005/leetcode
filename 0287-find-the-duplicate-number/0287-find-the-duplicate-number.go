func findDuplicate(nums []int) int {
    low, high := 0, len(nums) - 1
    for low < high {
        mid := (low + high) / 2
        leqMid := 0
        for _, n := range nums {
            if n <= mid {
                leqMid++
            }
        }
        
        if leqMid <= mid {
            low = mid + 1
        } else {
            high = mid
        }
    }
    
    return low
}