MAX_INT = 2 ** 31 - 1
MIN_INT = -1 * 2 ** 31 

VALID_CHARS = {
    "0": 0,
    "1": 1, 
    "2": 2, 
    "3": 3, 
    "4": 4, 
    "5": 5, 
    "6": 6, 
    "7": 7, 
    "8": 8,
    "9": 9,
}

class Solution:
    def myAtoi(self, s: str) -> int:
        while len(s) > 0 and s[0] == " ":
            s = s[1:]
                
        multiplier = 1        
        if len(s) > 0:
            if s[0] == "-":
                multiplier = -1
                s = s[1:]
            elif s[0] == "+":
                s = s[1:]
        
        res = 0
        while len(s) > 0 and s[0] in VALID_CHARS:
            res = (res * 10) + VALID_CHARS[s[0]]
            s = s[1:]
            
        res *= multiplier    
        
        if res < MIN_INT:
            return MIN_INT
            
        if MAX_INT < res:
            return MAX_INT
            
        return res 
