class Solution:
    def intToRoman(self, num: int) -> str:
        vals = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
        symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]
        
        roman = ""
        while num != 0:
            for i in range(len(vals)):
                if vals[i] <= num:
                    num -= vals[i]
                    roman +=  symbols[i]
                    break    
        
        return roman
                        
                        