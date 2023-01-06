func rotate(nums []int, k int)  {
    temp := make([]int, len(nums))
    for i := 0; i < len(temp); i++ {
        temp[(i+k)%len(temp)] = nums[i]
    }
    
    for i := 0; i < len(nums); i++ {
        nums[i] = temp[i]
    }
}

/*
func rotate(nums []int, k int)  {
    // reverse entire slice
    for i := 0; i < (len(nums)-1) / 2; i++ {
        nums[i], nums[len(nums)-1]
    }
}
*/