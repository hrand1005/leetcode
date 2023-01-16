class Solution:
    def addBinary(self, a: str, b: str) -> str:
        binsum = ""
        carry = 0
        while a != "" or b != "":
            digit_a, digit_b = 0, 0
            if a != "":
                digit_a = int(a[-1])
                a = a[:-1]
            if b != "":
                digit_b = int(b[-1])
                b = b[:-1]
                
            this_digit = digit_a + digit_b + carry    
            carry = 0
            if 1 < this_digit:
                this_digit -= 2
                carry = 1
            binsum = str(this_digit) + binsum
            
        if carry == 1:
            binsum = "1" + binsum
        return binsum
            
        