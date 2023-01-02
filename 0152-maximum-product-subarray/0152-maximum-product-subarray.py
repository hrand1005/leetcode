MIN_INT = -(1 << 31)

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        if self.product(nums) > 0:
            return self.product(nums)
        
        last_zero = -1
        candidates = []
        for i in range(len(nums)):
            if nums[i] == 0:
                candidates.append(nums[last_zero+1:i])
                last_zero = i
        candidates.append(nums[last_zero+1:])        
        
        max_product = MIN_INT
        if last_zero != -1:
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
    
    def product(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return MIN_INT
        
        product = 1
        for n in nums:
            product *= n
        return product    