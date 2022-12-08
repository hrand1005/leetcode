MAX_VAL = 2 ** 31 - 1
MIN_VAL = -1 * (2 ** 31)

class Solution:
    def reverse(self, x: int) -> int:
        multiplier = 1
        if x < 0:
            x *= -1
            multiplier = -1
        
        result = 0
        while x > 0:
            result = (result * 10) + (x % 10)
            if result > MAX_VAL or result < MIN_VAL:
                return 0
            
            x -= x % 10
            x /= 10
        
        return int(result) * multiplier

    """
class Solution:
    def reverse(self, x: int) -> int:
        max_val = 2 ** 31 - 1
        min_val = -1 * (2 ** 31)
        
        if x < 0:
               """ 