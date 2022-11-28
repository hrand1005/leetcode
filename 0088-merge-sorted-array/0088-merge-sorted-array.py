class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        if n == 0:
            return 
        
        write_index = m + n - 1
        i = m - 1
        j = n - 1
        while j >= 0:
            if i >= 0 and nums1[i] > nums2[j]:
                nums1[write_index] = nums1[i]
                i -= 1
            else:
                nums1[write_index] = nums2[j]
                j -= 1
            write_index -= 1
            