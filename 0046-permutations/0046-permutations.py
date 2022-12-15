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
            
            