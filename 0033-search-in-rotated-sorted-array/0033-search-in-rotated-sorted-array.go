/*
func recursiveSearch(nums []int, target, left, right int) int {
    if left > right {
        return -1
    }
    
    midpoint := (left + right) / 2
    if nums[midpoint] == target {
        return midpoint
    }
    
    if nums[left] <= nums[midpoint] {
        if nums[left] <= target && target < nums[midpoint] {
            return recursiveSearch(nums, target, left, midpoint - 1)
        } 
        return recursiveSearch(nums, target, midpoint + 1, right)
    } else {
        if nums[midpoint] < target && target <= nums[right] {
            return recursiveSearch(nums, target, midpoint + 1, right)
        }
        return recursiveSearch(nums, target, left, midpoint - 1)
    }
}

func search(nums []int, target int) int {
    return recursiveSearch(nums, target, 0, len(nums) - 1)
}
*/

func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    
    for left <= right {
        midpoint := (left + right) / 2
        if nums[midpoint] == target {
            return midpoint
        }
        
        if nums[left] <= nums[midpoint] {
            if nums[left] <= target && target < nums[midpoint] {
                right = midpoint - 1
            } else {
                left = midpoint + 1
            }
        } else {
            if nums[midpoint] < target && target <= nums[right] {
                left = midpoint + 1
            } else {
                right = midpoint - 1
            }
        }
    }
    
    return -1
}