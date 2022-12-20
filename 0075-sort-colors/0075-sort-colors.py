class Solution:
    def sortColors(self, nums: List[int]) -> None:
        left = 0
        for i in range(len(nums)):
            if nums[i] == 0:
                nums[i], nums[left] = nums[left], nums[i]
                left += 1
            
        right = len(nums) - 1
        i = 0
        while i <= right:    
            if nums[i] == 2:
                nums[i], nums[right] = nums[right], nums[i]
                right -= 1
            else:    
                i += 1    
                
        return        
        