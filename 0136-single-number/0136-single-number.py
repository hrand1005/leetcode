"""
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        occurs = {}
        for n in nums:
            if occurs.get(n) == None:
                occurs[n] = 1
            else:   
                occurs[n] += 1
        
        for k, v in occurs.items():
            if v == 1:
                return k
            
        return None
"""
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        res = 0
        for n in nums:
            res ^= n
            
        return res    