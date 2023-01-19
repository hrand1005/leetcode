import (
    "sort"
)

func wiggleSort(nums []int)  {
    sorted := make([]int, len(nums))
    copy(sorted, nums)
    sort.Ints(sorted)
    
    left := (len(sorted) - 1) / 2
    right := len(sorted) - 1
    
    for i := 0; i < len(sorted); i++ {
        if i % 2 == 0 {
            nums[i] = sorted[left]
            left--
        } else {
            nums[i] = sorted[right]
            right--
        }
    }
}