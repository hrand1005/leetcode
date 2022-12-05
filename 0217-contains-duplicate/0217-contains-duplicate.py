"""
class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        occurred = {}
        for n in nums:
            if occurred.get(n) != None:
                return True
            
            occurred[n] = True
        
        return False
"""

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        return len(nums) != len(set(nums))