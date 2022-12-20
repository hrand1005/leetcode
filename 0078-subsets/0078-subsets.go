/*
func subsets(nums []int) [][]int {
    return append([][]int{{}}, recursiveSubsets(nums)...)
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

func subsets(nums []int) [][]int {
    return dfs(nums, []int{})
}

func dfs(nums, cur []int) [][]int {
    curCopy := make([]int, len(cur))
    copy(curCopy, cur)
    subs := [][]int{ curCopy }
    for i := 0; i < len(nums); i++ {
        newCur := append(cur, nums[i])
        subs = append(subs, dfs(nums[i+1:], newCur)...)
    }
    return subs
}
*/

func subsets(nums []int) [][]int {
    all := [][]int{{}}
    for _, n := range nums {
        for _, s := range all {
            sCopy := make([]int, len(s))
            copy(sCopy, s)
            sub := append(sCopy, n)
            all = append(all, sub)
        }
    }
    
    return all
}