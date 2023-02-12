class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        val_to_idx = {}
        for i in range(len(nums)):
            comp = target - nums[i]
            compIdx = val_to_idx.get(comp)
            if compIdx != None:
                return [compIdx, i]
            val_to_idx[nums[i]] = i
        return []    
