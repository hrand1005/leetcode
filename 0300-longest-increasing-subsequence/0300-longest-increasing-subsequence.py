"""
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        # let each element be idx, sequence length
        max_sub = 0
        stack = []
        for i in range(len(nums)):
            stack.append((i, 1))
            
        while stack:
            idx, seq_len = stack.pop()
            max_sub = max(max_sub, seq_len)
            for i in range(idx+1, len(nums)):
                if nums[idx] < nums[i]:
                    stack.append((i, seq_len+1))
                    
        return max_sub            

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        self.cache = {}
        max_sub = 0
        for i in range(len(nums)):
            max_sub = max(max_sub, self.longest_from(i, nums))
        return max_sub    
        
    def longest_from(self, idx, nums) -> int:
        if self.cache.get(idx) != None:
            return self.cache[idx]
        
        max_sub = 1
        starts = [i for i in range(idx+1, len(nums)) if nums[idx] < nums[i]]
        for s in starts:
            max_sub = max(max_sub, 1+self.longest_from(s, nums))

        self.cache[idx] = max_sub
        return max_sub
"""
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        table = [0] * len(nums)
        max_sub = 1
        for i in range(len(nums)-1, -1, -1):
            max_sub_i = 1
            for j in range(i, len(nums)):
                if nums[i] < nums[j]:
                    max_sub_i = max(max_sub_i, 1+table[j])
            table[i] = max_sub_i
            
        return max(table)            