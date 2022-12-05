/*
import (
    "sort"
)

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    
    return false
}
*/

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

