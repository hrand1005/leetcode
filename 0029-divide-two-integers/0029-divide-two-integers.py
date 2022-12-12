MAX_INT = 2147483647
MIN_INT = -2147483648

"""
class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == 0:
            return 0
        
        sign = 1
        if dividend < 0:
            dividend = 0 - dividend
            sign = 0 - sign
            
        if divisor < 0:
            divisor = 0 - divisor
            sign = 0 - sign
        
        quotient = 0
        count = divisor
        while count <= dividend:
            scale = 1
            while (count + count) <= dividend:
                count += count
                scale += scale
            
            overflow, result = self.check_overflow(sign, quotient, scale, dividend)
            if overflow:
                return result
                
            quotient += scale
            dividend -= count
            count = divisor
            
        if sign < 0:
            quotient = 0 - quotient
        
        return quotient
"""    
class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        # if dividend == 0:
        #    return 0
        
        sign = 1
        if dividend < 0:
            dividend = 0 - dividend
            sign = 0 - sign
            
        if divisor < 0:
            divisor = 0 - divisor
            sign = 0 - sign
        
        quotient = 0
        counter = divisor
        while counter <= dividend:
            scale = 1
            while (counter << 1) <= dividend:
                scale <<= 1
                counter <<= 1
                
            overflow, res = self.check_overflow(sign, quotient, scale, dividend)
            if overflow:
                return res
        
            dividend -= counter
            quotient += scale
            counter = divisor
        
        if sign < 0:
            quotient = 0 - quotient
        
        return quotient

    
    def check_overflow(self, sign: int, quotient: int, to_add: int, dividend: int) -> (bool, int):
        """
        Checks whether adding 'to_add' to the quotient will cause
        overflow. 
        
        Returns True and the max/min int if overflow will occur, else returns False, 0
        """
        if sign < 0:
            if quotient > 0 - MIN_INT - to_add:
                return True, MIN_INT
        else:
            if quotient > MAX_INT - to_add:
                return True, MAX_INT
        
        return False, 0