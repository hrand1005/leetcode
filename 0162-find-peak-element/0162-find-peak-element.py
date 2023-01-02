class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        return self.find_peak_recursive(nums, 0, len(nums)-1)
        
    def find_peak_recursive(self, nums: List[int], left: int, right: int) -> int:
        if right == left:
            return right
        
        midpoint = (right + left) // 2
        if self.is_peak(nums, midpoint):
            return midpoint
        
        if 0 < midpoint and nums[midpoint] < nums[midpoint-1]:
            return self.find_peak_recursive(nums, left, midpoint-1)
        
        return self.find_peak_recursive(nums, midpoint+1, right)
    
    def is_peak(self, nums: List[int], midpoint: int) -> bool:
        if midpoint == 0:
            return nums[midpoint+1] < nums[midpoint]
        if midpoint == len(nums)-1:
            return nums[midpoint-1] < nums[midpoint]
        return nums[midpoint-1] < nums[midpoint] and nums[midpoint+1] < nums[midpoint]