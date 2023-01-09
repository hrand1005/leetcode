/*
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
*/

func findKthLargest(nums []int, k int) int {
    return quickSelect(nums, k, 0, len(nums)-1)
}

func quickSelect(nums []int, k, left, right int) int {
    if k == len(nums[left:right+1]) {
        return min(nums[left:right+1])
    }
    
    partitionIdx := partition(nums, left, right)
    
    if (right+1-partitionIdx) == k {
        return nums[partitionIdx]
    }
    if k <= (right-partitionIdx) {
        return quickSelect(nums, k, partitionIdx+1, right)
    }
    return quickSelect(nums, k-(right+1-partitionIdx), left, partitionIdx-1)
}

func partition(nums []int, left, right int) int {
    l, r := left, right
    pivot := nums[left]
    l++
    
    for l <= r {
        for l < right+1 && nums[l] <= pivot {
            l++
        }
        for pivot < nums[r] {
            r--
        }
        if l < r {
            nums[l], nums[r] = nums[r], nums[l]
        }
    }
    
    nums[left], nums[l-1] = nums[l-1], nums[left]
    return l-1
}

func min(s []int) int {
    m := s[0]
    for _, v := range s[1:] {
        if v < m {
            m = v
        }
    }
    return m
}