/*
func lengthOfLIS(nums []int) int {
    cache := make(map[int]int, len(nums))
    
    var longestRecursive func(int, []int) int 
    longestRecursive = func(idx int, nums []int) int {
        if v, ok := cache[idx]; ok {
            return v
        }
        maxSub := 1
        for i := idx+1; i < len(nums); i++ {
            if nums[idx] < nums[i] {
                maxSub = max(maxSub, 1+longestRecursive(i, nums))
            }
        }
        cache[idx] = maxSub
        return maxSub
    }
    
    longest := 1
    for i := 0; i < len(nums); i++ {
        longest = max(longest, longestRecursive(i, nums))
    }
    return longest
}

*/
func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func lengthOfLIS(nums []int) int {
    table := make([]int, len(nums))
    for i := len(nums)-1; i >= 0; i-- {
        table[i] = 1
        for j := i+1; j < len(nums); j++ {
            if nums[i] < nums[j] {
                table[i] = max(table[i], 1+table[j])
            }
        }
    }
    
    longest := 0
    for _, v := range table {
        longest = max(longest, v)
    }
    return longest
}