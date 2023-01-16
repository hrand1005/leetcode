class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        return self.search_insert_recursive(nums, target, 0, len(nums)-1)
    
    def search_insert_recursive(self, nums: List[int], target: int, left: int, right: int) -> int:
        if right <= left:
            if nums[left] < target:
                return left+1
            return left
        
        midpoint = (left+right)//2
        if nums[midpoint] == target:
            return midpoint
        if nums[midpoint] < target:
            return self.search_insert_recursive(nums, target, midpoint+1, right)
        return self.search_insert_recursive(nums, target, left, midpoint-1)