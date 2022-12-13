"""
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        return self.recursive_search(nums, target, 0, len(nums)-1)
    
    def recursive_search(self, nums: List[int], target: int, left: int, right: int) -> int:
        if left >= right:
            if nums[left] == target:
                return left
            return -1
        
        midpoint = (left + right) // 2
        if nums[midpoint] == target:
            return midpoint
        
        if nums[0] <= nums[midpoint]:
            if nums[0] <= target < nums[midpoint]:
                return self.recursive_search(nums, target, left, midpoint - 1)
            else:
                return self.recursive_search(nums, target, midpoint + 1, right)
        else:    
            if nums[midpoint] < target <= nums[len(nums)-1]:
                return self.recursive_search(nums, target, midpoint + 1, right)
            else:
                return self.recursive_search(nums, target, left, midpoint - 1)
 """           

class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        
        while left <= right:
            midpoint = (left + right) // 2
            if nums[midpoint] == target:
                return midpoint
            
            if nums[left] <= nums[midpoint]:
                if nums[left] <= target <= nums[midpoint]:
                    right = midpoint - 1
                else:
                    left = midpoint + 1
            else:
                if nums[midpoint] <= target <= nums[right]:
                    left = midpoint + 1
                else:
                    right = midpoint - 1
            
        return -1    