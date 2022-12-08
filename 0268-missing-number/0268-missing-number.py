"""
class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        occurred = {}

        for n in nums:
            occurred[n] = True

        for i in range(len(nums)+1):
            if not occurred.get(i):
                return i

        return None        
"""

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        sum_n = ((0 + len(nums)) / 2) * (len(nums) + 1)
        return int(sum_n) - sum(nums)
