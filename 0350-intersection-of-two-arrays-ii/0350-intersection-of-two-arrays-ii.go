/*
func intersect(nums1 []int, nums2 []int) []int {
    occurrences := map[int]int{}

    for _, v := range nums1 {
        occurrences[v]++
    }

    intersect := make([]int, 0, len(nums2))

    for _, v := range nums2 {
        if occ, ok := occurrences[v]; ok && occ > 0 {
            intersect = append(intersect, v)
            occurrences[v]--
        }
    }

    return intersect
}
*/

import (
    "sort"
)

func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)

    intersect := make([]int, 0, len(nums1))

    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            i++
        } else if nums2[j] < nums1[i] {
            j++
        } else {
            intersect = append(intersect, nums1[i])
            i++
            j++
        }
    }

    return intersect
}