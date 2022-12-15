"""
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        perms = {
            1: [[n] for n in nums]
        }
        
        for i in range(1, len(nums)+1):
            for j in range(i+1, len(nums)+1):
                last = perms[i]
                perms[j] = []
                
                for p in last:
                    for n in nums:
                        if n not in p:
                            perms[j].append(p + [n])
                    
        return perms[len(nums)]            

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        stack = [([], nums)]
        result = []
        
        while stack:
            perms, remaining = stack.pop()
            if not remaining:
                result.append(perms)
            for i in range(len(remaining)):
                new_perms = perms + [remaining[i]]
                new_remaining = remaining[:i] + remaining[i+1:]
                stack.append((new_perms, new_remaining))
                
        return result
"""            

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        if len(nums) == 1:
            return [[nums[0]]]
        
        result = []
        for i in range(len(nums)):
            this = nums[:i] + nums[i+1:]
            perms = self.permute(this)
            for p in perms:
                result.append([nums[i]] + p)
        
        return result