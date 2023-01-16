class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        if len(nums) == 1:
            if nums[0] == val:
                return 0
            return 1
        
        left, right = 0, len(nums)-1
        while left < right:
            while 0 < right and nums[right] == val:
                right -= 1
            while left < len(nums) and nums[left] != val:
                left += 1
            if left < right:
                nums[left], nums[right] = nums[right], nums[left]
                
        print(left)        
        return left