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
                if i % 2 == 0 and i < len(vals)-2:
                    if vals[i] - vals[i+2] <= num:
                        num -= vals[i] - vals[i+2]
                        roman += symbols[i+2] + symbols[i]
                        break
                if i % 2 == 1 and i < len(vals)-1:
                    if vals[i] - vals[i+1] <= num:
                        num -= vals[i] - vals[i+1]
                        roman += symbols[i+1] + symbols[i]
                        break
                   
        return roman