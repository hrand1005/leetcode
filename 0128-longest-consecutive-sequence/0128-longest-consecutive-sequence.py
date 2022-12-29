class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        num_set = set(nums)
        longest = 0
        
        for n in num_set:
            if n-1 not in num_set:
                consecutive = 1
                while n + consecutive in num_set:
                    consecutive += 1
                longest = max(longest, consecutive)    
        
        return longest