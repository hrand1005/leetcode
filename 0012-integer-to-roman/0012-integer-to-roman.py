"""
class Solution:
    def intToRoman(self, num: int) -> str:
        vals = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
        symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]
        roman = ""
        while num != 0:
            for i in range(len(vals)):
                if vals[i] <= num:
                    num -= vals[i]
                    roman += symbols[i]
                    break    
        return roman
"""

class Solution:
    def intToRoman(self, num: int) -> str:
        vals = [1000, 500, 100, 50, 10, 5, 1]
        symbols = ["M", "D", "C", "L", "X", "V", "I"]
        
        roman = ""
        while num != 0:
            for i in range(len(vals)):
                if vals[i] <= num:
                    num -= vals[i]
                    roman += symbols[i]
                    break
                offset = 2 - (i % 2)
                if i < len(vals)-offset:
                    if vals[i] - vals[i+offset] <= num:
                        num -= vals[i] - vals[i+offset]
                        roman += symbols[i+offset] + symbols[i]
                        break
                   
        return roman