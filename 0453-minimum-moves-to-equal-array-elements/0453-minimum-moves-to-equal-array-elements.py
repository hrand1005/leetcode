"""
class Solution:
    def minMoves(self, nums: List[int]) -> int:
        if len(nums) <= 1:
            return 0
        count = 0
        while not self.all_equal(nums):
            nums.sort()
            for i in range(len(nums)-1):
                nums[i] += 1
            count += 1    
        return count    
                
    def all_equal(self, nums: List[int]) -> bool:
        for i in range(1, len(nums)):
            if nums[i] != nums[i-1]:
                return False
        return True    
"""

class Solution:
    def minMoves(self, nums: List[int]) -> int:
        return sum(nums) - len(nums) * min(nums)
        