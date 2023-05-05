/*
func productExceptSelf(nums []int) []int {
    forwards := make([]int, len(nums))
    forwards[0] = nums[0]
    for i := 1; i < len(forwards); i++ {
        forwards[i] = nums[i] * forwards[i-1]
    }
    
    backwards := make([]int, len(nums))
    backwards[len(backwards)-1] = nums[len(nums)-1]
    for i := len(backwards)-2; i >= 0; i-- {
        backwards[i] = nums[i] * backwards[i+1]
    }
    
    result := make([]int, len(nums))
    for i := 1; i < len(result)-1; i++ {
        result[i] = forwards[i-1] * backwards[i+1]
    }
    result[0] = backwards[1]
    result[len(result)-1] = forwards[len(forwards)-2]
    
    return result
}
*/
func productExceptSelf(nums []int) []int {
    products := make([]int, len(nums))
    pbefore := 1
    for i := 0; i < len(products); i++ {
        products[i] = pbefore
        pbefore *= nums[i]
    }
    
    pafter := 1
    for i := len(nums)-1; i >= 0; i-- {
        products[i] *= pafter
        pafter *= nums[i]
    }
    return products
}