"""
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

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        local_min = nums[0]
        local_max = nums[0]
        max_product = nums[0]
        
        for i in range(1, len(nums)):
            prev_local_max = local_max
            prev_local_min = local_min
            local_max = max(max(prev_local_max*nums[i], nums[i]), prev_local_min*nums[i])
            local_min = min(min(prev_local_max*nums[i], nums[i]), prev_local_min*nums[i])
            max_product = max(max_product, local_max)
        
        return max_product
"""

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        max_product = nums[0]
        cur_product = 1
        for i in range(len(nums)):
            cur_product *= nums[i]
            max_product = max(max_product, cur_product)
            if cur_product == 0:
                cur_product = 1
                
        cur_product = 1
        for i in range(len(nums)-1, -1, -1):
            cur_product *= nums[i]
            max_product = max(max_product, cur_product)
            if cur_product == 0:
                cur_product = 1
        
        return max_product