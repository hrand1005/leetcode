class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        low_index, high_index = -1, -1
        
        left, right = 0, len(nums) - 1
        while left <= right:
            midpoint = (left + right) // 2
            
            if nums[midpoint] == target:
                low_index, high_index = self.find_lowest_and_highest(midpoint, target, nums)
                break
            
            if nums[midpoint] < target:
                left = midpoint + 1
            else:
                right = midpoint - 1
        
        
        return [low_index, high_index]
    
    def find_lowest_and_highest(self, midpoint: int, target: int, nums: List[int]) -> (int, int):
        low = midpoint
        while low - 1 >= 0 and nums[low-1] == target: 
            low -= 1
        
        high = midpoint
        while high + 1 <= len(nums) - 1 and nums[high+1] == target:
            high += 1
        
        return low, high
        