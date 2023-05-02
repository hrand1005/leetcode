func twoSum(nums []int, target int) []int {
    comps := make(map[int]int, len(nums))
    for i, n := range nums {
        if v, ok := comps[target-n]; ok {
            return []int{v, i}
        }
        comps[n] = i
    }
    return nil
}