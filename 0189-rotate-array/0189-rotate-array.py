class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        tmp = [0] * len(nums)
        for i in range(n):
            tmp[(i+k)%n] = nums[i]
        
        for i in range(n):
            nums[i] = tmp[i]
        
        return