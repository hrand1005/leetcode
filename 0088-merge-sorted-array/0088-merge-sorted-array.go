func merge(nums1 []int, m int, nums2 []int, n int)  {
    writeTo := m + n - 1
    m--
    n--
    
    for n >= 0 {
        if m >= 0 && nums1[m] > nums2[n] {
            nums1[writeTo] = nums1[m]
            m--
        } else {
            nums1[writeTo] = nums2[n]
            n--
        }
        writeTo--
    }
}
