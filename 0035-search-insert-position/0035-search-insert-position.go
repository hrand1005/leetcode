func searchInsert(nums []int, target int) int {
    return searchInsertRecursive(nums, target, 0, len(nums)-1)
}

func searchInsertRecursive(nums []int, target, low, high int) int {
    if high <= low {
        if nums[low] < target {
            return low+1
        }
        return low
    }
    midpoint := (low+high)/2
    if nums[midpoint] == target {
        return midpoint
    }
    if nums[midpoint] < target {
        return searchInsertRecursive(nums, target, midpoint+1, high)
    }
    return searchInsertRecursive(nums, target, 0, midpoint-1)
}