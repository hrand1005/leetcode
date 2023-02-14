class Solution:
    def minStartValue(self, nums: List[int]) -> int:
        min_val = nums[0]
        total = nums[0]
        for n in nums[1:]:
            total += n
            min_val = min(min_val, total)
        return 1 - min(min_val, 0)
        
        