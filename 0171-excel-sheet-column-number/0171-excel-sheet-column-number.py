class Solution(object):
    def titleToNumber(self, columnTitle):
        col_val = {}
        alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        for i in range(26):
            col_val[alphabet[i]] = i + 1
        
        total = col_val[columnTitle[0]]
        for c in columnTitle[1:]:
            total = total * 26 + col_val[c]
        
        return total
            
            
        