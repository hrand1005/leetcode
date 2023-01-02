class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        candidates = self.find_candidates(nums)
        max_product = nums[0]
        if len(candidates) > 1:
            max_product = 0
        
        while candidates:
            c = candidates.pop()
            prod = self.product(c)
            if prod < 0:
                for i in range(len(c)):
                    if c[i] < 0:
                        candidates.append(c[:i])
                        candidates.append(c[i+1:])
            max_product = max(max_product, prod)            
            
        return max_product
    
    # finds candidate subarrays for max product
    def find_candidates(self, nums: List[int]) -> List[int]:
        last_zero = -1
        candidates = []
        for i in range(len(nums)):
            if nums[i] == 0:
                candidates.append(nums[last_zero+1:i])
                last_zero = i
        candidates.append(nums[last_zero+1:])        
        return candidates
    
    # returns min int32 val if nums is empty
    def product(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return -(1 << 31)
        product = 1
        for n in nums:
            product *= n
        return product    