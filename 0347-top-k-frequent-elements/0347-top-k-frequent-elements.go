func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int, len(nums))
    for _, n := range nums {
        count[n]++
    }
    
    inverse := make(map[int][]int, len(count))
    for k, v := range count {
        if s, ok := inverse[v]; ok {
            inverse[v] = append(s, k)
        } else {
            inverse[v] = []int{k}
        }
    }
    
    result := make([]int, 0, k)
    i := len(nums)
    for len(result) < k {
        if v, ok := inverse[i]; ok {
            result = append(result, v...)
        }
        i--
    }
    
    return result
}