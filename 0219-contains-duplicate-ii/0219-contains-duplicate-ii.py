class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        seen = {}
        for i in range(len(nums)):
            j = seen.get(nums[i])
            if j != None and abs(i-j) <= k:
                return True
            seen[nums[i]] = i
        return False    
                