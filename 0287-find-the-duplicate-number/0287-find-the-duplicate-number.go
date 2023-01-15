/*
func findDuplicate(nums []int) int {
    low, high := 0, len(nums) - 1
    for low < high {
        mid := (low + high) / 2
        leqMid := 0
        for _, n := range nums {
            if n <= mid {
                leqMid++
            }
        }
        
        if leqMid <= mid {
            low = mid + 1
        } else {
            high = mid
        }
    }
    
    return low
}
*/

func findDuplicate(nums []int) int {
    slow, fast := nums[0], nums[nums[0]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    
    start := 0
    for start != slow {
        start = nums[start]
        slow = nums[slow]
    }
    
    return start
}