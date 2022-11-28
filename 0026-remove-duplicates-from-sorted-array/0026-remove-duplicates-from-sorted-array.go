func removeDuplicates(nums []int) int {
    i := 1
    for i < len(nums) {
        if nums[i] == nums[i-1] {
            nums = append(nums[:i], nums[i+1:]...)
        } else {
            i++
        }
    }
    return i
}

/*
func removeDuplicates(nums []int) int {
    last := nums[0]
    insertAt := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != last {
            last = nums[i]
            nums[insertAt] = last
            insertAt++
        } 
    }
    
    return insertAt
} 
*/