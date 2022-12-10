NUM_MAP = {
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

"""
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if len(digits) == 0:
            return []
        
        combos = {}
        combos[0] = list(NUM_MAP[digits[0]])
        for i in range(1, len(digits)):
            combos[i] = []
            for combo in combos[i-1]:
                for letter in list(NUM_MAP[digits[i]]):
                    combos[i].append(combo + letter)
                    
        return combos[len(digits)-1]
"""

class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if len(digits) == 0:
            return []
        
        this = list(NUM_MAP[digits[-1]])
        
        if len(digits) == 1:
            return this
        
        combos = []
        prev = self.letterCombinations(digits[:-1])
        for letter in prev:
            for c in this:
                combos.append(letter + c)
        
        return combos