/*
import (
    "math"
)

func jump(nums []int) int {
    cache := make(map[int]int, len(nums))
    
    var jumpRec func([]int, int) int
    jumpRec = func(nums []int, idx int) int {
        if v, ok := cache[idx]; ok {
            return v
        }
        if idx == len(nums)-1 {
            return 0
        }
        scope := nums[idx]
        if idx + scope >= len(nums)-1 {
            return 1
        }
        minJumps := math.MaxInt32
        for i := idx+1; i <= idx + scope; i++ {
            minJumps = min(minJumps, jumpRec(nums, i))
        }
        minJumps++
        cache[idx] = minJumps
        return minJumps
    }
    
    return jumpRec(nums, 0)
}
*/

func jump(nums []int) int {
    table := make([]int, len(nums))
    for i := 0; i < len(table); i++ {
        table[i] = math.MaxInt32
    }
    table[0] = 0
    
    for i := 0; i < len(nums); i++ {
        for j := i+1; j <= i+nums[i]; j++ {
            if j < len(nums) {
                table[j] = min(table[j], table[i]+1)
            }
        }
    }
    
    return table[len(table)-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


/*
func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    start, end := 0, nums[0]
    jumps := 1
    for end < len(nums)-1 {
        jumps++
        scope := 0
        for i := start; i < end+1; i++ {
            scope = max(scope, i+nums[i])
        }
        start, end = end, scope
    }
    return jumps
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
*/