class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        for n in nums:
            count[n] = count.get(n, 0) + 1
        
        inverse = {}
        for key, val in count.items():
            inverse[val] = inverse.get(val, []) + [key]
        
        result = []
        i = len(nums)
        while len(result) < k:
            if inverse.get(i) is not None:
                result += inverse[i]
            i -= 1    
        
        return result