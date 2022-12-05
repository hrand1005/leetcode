func containsDuplicate(nums []int) bool {
    count := map[int]int{}
    
    for _, v := range nums {
        if _, ok := count[v]; ok {
            return true
        }
        count[v]++
    }
    
    return false
}