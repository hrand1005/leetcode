"""
POS_MAP = { 
    'A': 1,
    'B': 2,
    'C': 3,
    'D': 4,
    'E': 5,
    'F': 6,
    'G': 7,
    'H': 8,
    'I': 9,
    'J': 10,
    'K': 11,
    'L': 12,
    'M': 13,
    'N': 14,
    'O': 15,
    'P': 16,
    'Q': 17,
    'R': 18,
    'S': 19,
    'T': 20,
    'U': 21,
    'V': 22,
    'W': 23,
    'X': 24,
    'Y': 25,
    'Z': 26,
}

class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        return self.recursive(0, columnTitle)
    
    def recursive(self, multiplier, columnTitle) -> int:
        if len(columnTitle) == 0:
            return 0
        
        last = POS_MAP[columnTitle[len(columnTitle)-1]]
        return 26 ** multiplier * last + self.recursive(multiplier + 1, columnTitle[:-1])


class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        col = 0
        i = len(columnTitle) - 1
        j = 0
        while i >= 0:
            last_char = columnTitle[i]
            col += 26 ** j * POS_MAP[last_char]
            i -= 1
            j += 1
        
        return col
        
"""

class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        col = 0
        i = len(columnTitle) - 1
        j = 0
        while i >= 0:
            last_char = columnTitle[i]
            col += 26 ** j * (ord(last_char) - 64)
            i -= 1
            j += 1
        
        return col
        
