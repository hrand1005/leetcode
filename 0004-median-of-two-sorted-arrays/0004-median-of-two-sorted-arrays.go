func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {
        nums1, nums2 = nums2, nums1
    }
    fmt.Println(nums1)
    len1, len2 := len(nums1), len(nums2)
    full := len1 + len2
    
    low, high := 0, len1
    for low <= high {
        mid1 := (low + high) / 2
        mid2 := full / 2 - mid1
        
        left1, right1 := math.MinInt32, math.MaxInt32
        if 0 < mid1 {
            left1 = nums1[mid1-1]
        }
        if mid1 < len1 {
            right1 = nums1[mid1]
        }
        
        left2, right2 := math.MinInt32, math.MaxInt32
        if 0 < mid2 {
            left2 = nums2[mid2-1]
        }
        if mid2 < len2 {
            right2 = nums2[mid2]
        }
        
        if med, ok := median(left1, right1, left2, right2, full); ok {
            return med
        } else if right2 < left1 {
            high = mid1 - 1
        } else {
            low = mid1 + 1
        }
    }
    
    return 0
}

func median(l1, r1, l2, r2, full int) (float64, bool) {
    if l1 <= r2 && l2 <= r1 {
        if full % 2 == 0 {
            return float64((max(l1, l2) + min(r1, r2))) / 2, true
        }
        return float64(min(r1, r2)), true
    }
    return 0, false
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}