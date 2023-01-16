"""
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        if len(nums) == 1:
            if nums[0] == val:
                return 0
            return 1
        
        left, right = 0, len(nums)-1
        while left < right:
            while 0 < right and nums[right] == val:
                right -= 1
            while left < len(nums) and nums[left] != val:
                left += 1
            if left < right:
                nums[left], nums[right] = nums[right], nums[left]
                
        return left

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        start, end = 0, len(nums) - 1
        while start <= end:
            if nums[start] == val:
                nums[start], nums[end] = nums[end], nums[start]
                end -= 1
            else:
                start += 1
        return start        
        
"""

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        i = 0
        for n in nums:
            if n != val:
                nums[i] = n
                i += 1    
        return i        