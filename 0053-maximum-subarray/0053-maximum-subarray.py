"""
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        count = 0
        max_sum = nums[0]
        for i in range(len(nums)):
            if count + nums[i] < 0:
                if nums[i] > max_sum:
                    max_sum = nums[i]
                count = 0
            else:
                count += nums[i]
                if count > max_sum:
                    max_sum = count
        
        return max_sum
"""

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        sums = [0] * len(nums)
        sums[0] = nums[0]
        
        for i in range(1, len(nums)):
            sums[i] = max(sums[i-1] + nums[i], nums[i])
        
        return max(sums)