class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        perms = {
            1: []
        }
        
        # at perms[1]
        for n in nums:
            perms[1].append([n])

            
        for i in range(1, len(nums)+1):
            for j in range(i+1, len(nums)+1):
                last = perms[i]
                perms[j] = []
                
                for n in nums:
                    for p in last:
                        if n not in p:
                            new = p + [n]
                            perms[j].append(p + [n])
                    
        return perms[len(nums)]            
            
            