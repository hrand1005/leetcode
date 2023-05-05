class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        products = [nums[0]]
        for i in range(1, len(nums)):
            products.append(products[i-1] * nums[i])
        
        reverse = 1
        products[-1] = products[-2]
        for i in range(len(nums)-2, 0, -1):
            reverse *= nums[i+1]
            products[i] = products[i-1] * reverse
        products[0] = reverse * nums[1]
        return products
