func sortColors(nums []int)  {
    left, right := 0, len(nums) - 1
    idx := 0
    for idx <= right {
        if nums[idx] == 0 {
            nums[idx], nums[left] = nums[left], nums[idx]
            left++
            idx++
        } else if nums[idx] == 2 {
            nums[idx], nums[right] = nums[right], nums[idx]
            right--
        } else {
            idx++
        }
    }
}