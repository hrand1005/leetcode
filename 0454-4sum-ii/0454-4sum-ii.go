func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    oneTwoSums := map[int]int{}
    for _, one := range nums1 {
        for _, two := range nums2 {
            oneTwoSums[one+two]++
        }
    }
    
    tuples := 0
    for _, three := range nums3 {
        for _, four := range nums4 {
            tuples += oneTwoSums[-(three+four)]
        }
    }
    
    return tuples
}