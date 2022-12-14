func searchRange(nums []int, target int) []int {
    return []int{
        findFirstOccurrence(nums, target),
        findLastOccurrence(nums, target),
    }
}

func findFirstOccurrence(nums []int, target int) int {
    firstIndex := -1
    
    left, right := 0, len(nums) - 1
    for left <= right {
        midpoint := (left + right) / 2
        
        if nums[midpoint] == target {
            firstIndex = midpoint
        }
        
        if nums[midpoint] >= target {
            right = midpoint - 1
        } else {
            left = midpoint + 1
        }
    }
    
    return firstIndex
}

func findLastOccurrence(nums []int, target int) int {
    lastIndex := -1
    
    left, right := 0, len(nums) - 1
    for left <= right {
        midpoint := (left + right) / 2
        
        if nums[midpoint] == target {
            lastIndex = midpoint
        }
        
        if nums[midpoint] <= target {
            left = midpoint + 1
        } else {
            right = midpoint - 1
        }
    }
    
    return lastIndex
}