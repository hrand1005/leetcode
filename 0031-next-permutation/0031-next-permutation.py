class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        i = len(nums)-1
        while 0 < i and nums[i] <= nums[i-1]:
            i -= 1
        if i == 0:
            nums.reverse()
            return
        k = i - 1
        j = len(nums)-1
        while nums[j] <= nums[k]:
            j -= 1
            
        nums[j], nums[k] = nums[k], nums[j]
        l, r = k+1, len(nums)-1
        
        while l < r:
            nums[l], nums[r] = nums[r], nums[l]
            l += 1
            r -= 1
        return     
        
