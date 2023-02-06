"""
class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()
        quadruplets = set()
        result = []
        for i in range(len(nums)-3):
            for j in range(i+1, len(nums)):
                for k in range(j+1, len(nums)):
                    for l in range(k+1, len(nums)):
                        quad = [nums[i], nums[j], nums[k], nums[l]]
                        this = nums[i] + nums[j] + nums[k] + nums[l]
                        if this == target and tuple(quad) not in quadruplets:
                            quadruplets.add(tuple(quad))
                            result.append(quad)
        return result                   
"""

class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()
        quads = set()
        
        subproblems = {}
        for i in range(len(nums)):
            for j in range(i+3, len(nums)):
                subproblems[(i,j)] = nums[i] + nums[j]
        
        for k, v in subproblems.items():
            start, end = k[0]+1, k[1]
            sums = self.two_sum(nums[start:end], target-v)
            for s in sums:
                quads.add((nums[k[0]], s[0], s[1], nums[k[1]]))
        
        result = []
        for q in quads:
            result.append(list(q))
        return result    
        
    def two_sum(self, nums: List[int], target: int) -> Set[Tuple[int]]:
        res = set()
        occ = {}
        for i in range(len(nums)):
            comp = target - nums[i]
            if occ.get(comp) != None:
                res.add((comp, nums[i]))
            occ[nums[i]] = i
        return res    