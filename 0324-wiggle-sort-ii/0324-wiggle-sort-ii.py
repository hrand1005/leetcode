class Solution:
    def wiggleSort(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        increasing = nums.copy()
        increasing.sort()
        left = (len(nums)-1)//2
        right = len(nums)-1
        for i in range(len(nums)):
            if i%2 == 0:
                nums[i] = increasing[left]
                left -= 1
            else:
                nums[i] = increasing[right]
                right -= 1
