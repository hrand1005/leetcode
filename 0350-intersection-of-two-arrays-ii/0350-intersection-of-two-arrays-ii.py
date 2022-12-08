"""
class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        occurrences = {}

        for n in nums1:
            if not occurrences.get(n):
                occurrences[n] = 1
            else:
                occurrences[n] += 1

        intersect = []

        for n in nums2:
            if occurrences.get(n, 0) > 0:
                intersect.append(n)
                occurrences[n] -= 1

        return intersect
"""

class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        intersect = []

        nums1.sort()
        nums2.sort()

        i, j = 0, 0
        while i < len(nums1) and j < len(nums2):
            if nums1[i] < nums2[j]:
                i += 1
            elif nums2[j] < nums1[i]:
                j += 1
            else:
                intersect.append(nums1[i])    
                i += 1
                j += 1

        return intersect        