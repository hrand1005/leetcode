class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        s = set(nums)
        longest = 0
        for n in nums:
            if n-1 not in s:
                seq = 1
                while n+seq in s:
                    seq += 1
                longest = max(longest, seq)    
        return longest