func findDifference(nums1 []int, nums2 []int) [][]int {
    map1 := make(map[int]bool)
    for i := 0; i < len(nums1); i++ {
        map1[nums1[i]] = true
    }
    
    map2 := make(map[int]bool)
    for i := 0; i < len(nums2); i++ {
        if _, ok := map1[nums2[i]]; ok {
            map1[nums2[i]] = false
        } else {
            map2[nums2[i]] = true
        }
    }
    
    unique1 := make([]int, 0, len(nums1))
    for k, v := range map1 {
        if v {
            unique1 = append(unique1, k)
        }
    }
    unique2 := make([]int, 0, len(nums2))
    for k, v := range map2 {
        if v {
            unique2 = append(unique2, k)
        }
    }
    return [][]int{unique1, unique2}
}