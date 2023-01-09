func findKthLargest(nums []int, k int) int {
    lessThan := make([]int, 0, len(nums))
    equalTo := make([]int, 0, len(nums))
    greaterThan := make([]int, 0, len(nums))
    
    pivot := nums[0]
    for _, n := range nums {
        if n < pivot {
            lessThan = append(lessThan, n)
        } else if n == pivot {
            equalTo = append(equalTo, n)
        } else {
            greaterThan = append(greaterThan, n)
        }
    }
    
    if k <= len(greaterThan) {
        return findKthLargest(greaterThan, k)
    }
    if len(equalTo) + len(greaterThan) < k {
        return findKthLargest(lessThan, k-len(equalTo)-len(greaterThan))
    }
    return equalTo[0]
}