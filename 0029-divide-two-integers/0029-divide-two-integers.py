MAX_INT = 2**31 - 1
MIN_INT = 2**31 * -1

class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == 0:
            return 0
        
        multiplier = 1
        if dividend < 0:
            dividend = 0 - dividend
            multiplier = 0 - multiplier
            
        if divisor < 0:
            divisor = 0 - divisor
            multiplier = 0 - multiplier
        
        quotient = 0
        count = divisor
        while count <= dividend:
            scale = 1
            while (count + count) <= dividend:
                count += count
                scale += scale
            
            if multiplier < 0:
                if quotient > 0 - MIN_INT - scale:
                    return MIN_INT
            else:
                if quotient > MAX_INT - scale:
                    return MAX_INT
                
            quotient += scale
            dividend -= count
            count = divisor
            
        if multiplier < 0:
            quotient = 0 - quotient
        
        return quotient