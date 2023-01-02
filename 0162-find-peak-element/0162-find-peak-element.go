func findPeakElement(nums []int) int {
    return findPeakRecursive(nums, 0, len(nums)-1)
}

func findPeakRecursive(nums []int, left, right int) int {
    if left == right {
        return right
    }
    
    midpoint := (left + right) / 2
    if isPeak(nums, midpoint) {
        return midpoint
    }
    
    if nums[midpoint] != 0 && nums[midpoint] < nums[midpoint+1] {
        return findPeakRecursive(nums, midpoint+1, right)
    }
    return findPeakRecursive(nums, left, midpoint-1)
}

func isPeak(nums []int, midpoint int) bool {
    if midpoint == 0 {
        return nums[midpoint+1] < nums[midpoint]
    }
    if midpoint == len(nums)-1 {
        return nums[midpoint-1] < nums[midpoint]
    }
    return nums[midpoint-1] < nums[midpoint] && nums[midpoint+1] < nums[midpoint]
}