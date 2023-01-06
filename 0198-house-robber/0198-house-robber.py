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