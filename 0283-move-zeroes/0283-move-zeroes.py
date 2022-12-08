"""
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        i, j = 0, 0
        while True:
            while nums[i] != 0:
                i += 1
                if i == len(nums):
                    return

            j = i

            while nums[j] == 0:
                j += 1
                if j == len(nums):
                    return

            nums[i], nums[j] = nums[j], nums[i]

            i += 1
"""

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        i, j = 0, 0
        while j < len(nums):
            if nums[i] == 0 and nums[j] != 0:
                nums[i], nums[j] = nums[j], nums[i]

            if nums[i] != 0:
                i += 1

            j += 1

