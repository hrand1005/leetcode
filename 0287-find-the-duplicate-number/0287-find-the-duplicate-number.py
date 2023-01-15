"""
class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            if nums[i] in nums[i+1:]:
                return nums[i]
        return 0       
        
class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        low, high = 0, len(nums)-1
        while low < high:
            midpoint = (low + high) // 2
            leq_count = 0
            for n in nums:
                if n <= midpoint:
                    leq_count += 1
                    
            if leq_count <= midpoint:
                low = midpoint + 1
            else:
                # NOTE: the midpoint itself could be the duplicate
                high = midpoint
        
        return low
"""

class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        slow, fast = nums[0], nums[nums[0]]
        while slow != fast:
            slow = nums[slow]
            fast = nums[nums[fast]]
        
        start = 0
        while start != slow:
            start = nums[start]
            slow = nums[slow]
        
        return start