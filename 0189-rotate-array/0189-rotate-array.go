/*
func rotate(nums []int, k int)  {
    temp := make([]int, len(nums))
    for i := 0; i < len(temp); i++ {
        temp[(i+k)%len(temp)] = nums[i]
    }
    
    for i := 0; i < len(nums); i++ {
        nums[i] = temp[i]
    }
}
*/

func rotate(nums []int, k int)  {
    k %= len(nums)
    reverse(nums, 0, len(nums)-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, left, right int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
