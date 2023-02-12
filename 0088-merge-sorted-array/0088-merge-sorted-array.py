class Solution(object):
    def merge(self, nums1, m, nums2, n):
        """
        :type nums1: List[int]
        :type m: int
        :type nums2: List[int]
        :type n: int
        :rtype: None Do not return anything, modify nums1 in-place instead.
        """
        idx = len(nums1)-1
        l, r = len(nums1)-len(nums2)-1, len(nums2)-1
        
        while idx >= 0:
            if l >= 0 and r >= 0:
                if nums1[l] > nums2[r]:
                    nums1[idx] = nums1[l]
                    l -= 1
                else:
                    nums1[idx] = nums2[r]
                    r -= 1
            elif l < 0:
                nums1[idx] = nums2[r]
                r -= 1
            else:
                nums1[idx] = nums1[l]
                l -= 1
            idx -= 1    
            
        return    
        