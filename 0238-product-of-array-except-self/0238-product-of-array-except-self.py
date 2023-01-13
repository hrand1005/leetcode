"""
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        forward = [0] * len(nums)
        forward[0] = nums[0]
        for i in range(1, len(nums)):
            forward[i] = forward[i-1] * nums[i]
        
        backward = [0] * len(nums)
        backward[-1] = nums[-1]
        for i in range(len(nums)-2, -1, -1):
            backward[i] = backward[i+1] * nums[i]
        
        result = [0] * len(nums)
        result[0] = backward[1]
        result[-1] = forward[-2]
        for i in range(1, len(nums)-1):
            result[i] = forward[i-1] * backward[i+1]
        
        return result
"""

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = [0] * len(nums)
        product = 1
        for i in range(0, len(nums)):
            result[i] = product * nums[i]
            product = result[i]
        
        product = 1
        for i in range(len(nums)-1, 0, -1):
            result[i] = product * result[i-1]
            product *= nums[i]
        result[0] = product
        
        return result