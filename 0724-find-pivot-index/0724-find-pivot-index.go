func pivotIndex(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    
    acc := 0
    for i, n := range nums {
        if (total - n - acc) == acc {
            return i
        }
        acc += n
    }
    return -1
}