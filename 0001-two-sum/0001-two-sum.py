class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        comps = {}
        for i, n in enumerate(nums):
            if target-n in comps:
                return [comps[target-n], i]
            comps[n] = i
        return []    