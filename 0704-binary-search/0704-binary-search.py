class Solution:
    def search(self, nums: List[int], target: int) -> int:
        return self.binary_search(nums, target, 0, len(nums)-1)
    
    def binary_search(self, nums: List[int], target: int, lo: int, hi: int) -> int:
        if lo > hi:
            return -1
        
        mid = (lo + hi) // 2
        if nums[mid] == target:
            return mid
        
        if nums[mid] > target:
            return self.binary_search(nums, target, lo, mid-1)
        
        return self.binary_search(nums, target, mid+1, hi)
        
        