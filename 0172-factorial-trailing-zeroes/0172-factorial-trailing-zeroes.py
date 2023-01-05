class Solution:
    def trailingZeroes(self, n: int) -> int:
        if n == 0:
            return 0
        
        product = 1
        while n > 0:
            product *= n
            n -= 1
        
        str_product = str(product)
        zero_count = 0
        while str_product[-1] == "0":
            zero_count += 1
            str_product = str_product[:-1]
        
        return zero_count
        