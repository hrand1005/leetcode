func majorityElement(nums []int) int {
    maj := (len(nums) / 2) + 1
    occ := map[int]int{}
    
    for _, n := range nums {
        occ[n]++
        if occ[n] == maj {
            return n
        } 
    }
    return -1 // should never occur
}