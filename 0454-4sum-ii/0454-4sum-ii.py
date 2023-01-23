"""
class Solution:
    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        tuples = []
        for i in range(len(nums1)):
            for j in range(len(nums2)):
                for k in range(len(nums3)):
                    for l in range(len(nums4)):
                        if nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0:
                            tuples.append((i, j, k, l))
        return len(tuples)
"""

class Solution:
    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        one_two = {}
        for i in range(len(nums1)):
            for j in range(len(nums2)):
                sum_ij = nums1[i] + nums2[j]
                one_two[sum_ij] = one_two.get(sum_ij, 0) + 1
                
        tuples = 0        
        for i in range(len(nums3)):
            for j in range(len(nums4)):
                sum_ij = nums3[i] + nums4[j]
                tuples += one_two.get(-sum_ij, 0)
                
        return tuples        
                    
                    