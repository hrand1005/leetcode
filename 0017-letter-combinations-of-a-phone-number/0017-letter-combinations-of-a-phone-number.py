class Solution(object):
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        num_to_letter = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz",
        }
        
        if len(digits) == 0:
            return []
        if len(digits) == 1:
            return list(num_to_letter[digits])
        
        this = list(num_to_letter[digits[0]])
        rest = self.letterCombinations(digits[1:])
        combos = []
        for t in this:
            for r in rest:
                combos.append(t+r)
        return combos    