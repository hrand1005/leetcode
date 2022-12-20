func subsets(nums []int) [][]int {
    emptySet := []int{}
    return append([][]int{emptySet}, recursiveSubsets(nums)...)
}

func recursiveSubsets(nums []int) [][]int {
    if len(nums) == 0 {
        return nil
    }
    
    allSubsets := [][]int{}
    for i, v := range nums {
        subs := recursiveSubsets(nums[i+1:])
        for i := 0; i < len(subs); i++ {
            subs[i] = append([]int{v}, subs[i]...)
        }
        subs = append([][]int{{v}}, subs...)
        allSubsets = append(allSubsets, subs...)
    }
    
    return allSubsets
}