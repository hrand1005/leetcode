"""
class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        n = len(nums)
        tmp = [0] * len(nums)
        for i in range(n):
            tmp[(i+k)%n] = nums[i]
        
        for i in range(n):
            nums[i] = tmp[i]
        
        return
"""

class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        n = len(nums)
        new_loc = {}
        for i in range(n):
            new_loc[(i+k)%n] = nums[i]
        
        for i in range(n):
            nums[i] = new_loc[i]
        
        return