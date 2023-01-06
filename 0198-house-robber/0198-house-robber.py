"""
class Solution:
    def rob(self, nums: List[int]) -> int:
        self.cache = {}
        return self.rob_recursive(0, nums)
                         
    def rob_recursive(self, idx: int, nums: List[int]) -> int:
        if len(nums) <= idx:
            return 0
        if len(nums)-1 == idx:
            return nums[idx]
        
        if self.cache.get(idx) != None:
            return self.cache[idx]
        
        prof_this = nums[idx] + self.rob_recursive(idx+2, nums)
        prof_next = nums[idx+1] + self.rob_recursive(idx+3, nums)
        max_profit = max(prof_this, prof_next)
        self.cache[idx] = max_profit
        
        return max_profit                 
"""

class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        
        table = [0] * len(nums)
        table[0] = nums[0]
        table[1] = max(nums[1], nums[0])
        
        for i in range(2, len(nums)):
            prof_this = nums[i] + table[i-2]
            prof_prev = table[i-1]
            table[i] = max(prof_this, prof_prev)
        
        return table[-1]
            
            