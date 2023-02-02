INT_MIN = -(2 ** 32)
INT_MAX = 2 ** 32

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        n1 = len(nums1)
        n2 = len(nums2)
        
        if n2 < n1:
            nums1, nums2 = nums2, nums1
            n1, n2 = n2, n1
        
        low, high = 0, n1
        while low <= high:
            mid1 = (low + high) // 2
            mid2 = (n1 + n2) // 2 - mid1

            l1, r1 = INT_MIN, INT_MAX
            if 0 < mid1:
                l1 = nums1[mid1-1]
            if mid1 < n1:
                r1 = nums1[mid1]
            
            l2, r2 = INT_MIN, INT_MAX
            if 0 < mid2:
                l2 = nums2[mid2-1]
            if mid2 < n2:
                r2 = nums2[mid2]
            
            if l1 <= r2 and l2 <= r1:
                if (n1 + n2) % 2 == 0:
                    return (max(l1, l2) + min(r1, r2)) / 2
                else:
                    return min(r1, r2)
            elif r2 < l1:
                high = mid1 - 1
            else:
                low = mid1 + 1
                
        return 0.0        