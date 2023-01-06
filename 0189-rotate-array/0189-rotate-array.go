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
    k = k % len(nums)
    if k == 0 || len(nums) <= 1 {
        return
    }
    
    n := len(nums)-1
    for i := 0; i <= n/2; i++ {
        nums[i], nums[n-i] = nums[n-i], nums[i]
    }
    
    first := k-1
    for i := 0; i <= first/2; i++ {
        nums[i], nums[first-i] = nums[first-i], nums[i]
    }
    
    for i := 0; i <= (n-k)/2; i++{
        nums[k+i], nums[n-i] = nums[n-i], nums[k+i]
    }
}