class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        comps = {}
        for i, n in enumerate(nums):
            if comps.get(target-n) is not None:
                return [comps.get(target-n), i]
            comps[n] = i
        return []    