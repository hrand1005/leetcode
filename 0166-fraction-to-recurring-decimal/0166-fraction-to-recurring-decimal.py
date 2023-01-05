class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        sign = ""
        if numerator < 0 < denominator or denominator < 0 < numerator:
            sign += "-"
        
        numerator, denominator = abs(numerator), abs(denominator)
        whole = numerator // denominator
        rem = numerator % denominator
        if rem == 0:
            return sign + str(whole)
        
        dec = ""
        seen = {
            rem: 0
        }
        n = rem * 10
        while rem != 0:
            dec += str(n // denominator)
            rem = n % denominator
            if seen.get(rem) != None:
                break
            seen[rem] = len(dec)
            n = rem * 10
            
        if rem != 0:    
            start = seen.get(rem)
            dec = f"{dec[:start]}({dec[start:]})"
            
        return f"{sign}{whole}.{dec}"