class Solution:
    def isValid(self, s: str) -> bool:
        pair = {
            '(': ')',
            '{': '}',
            '[': ']'
        }
        
        stack = []
        for char in s:
            if char in pair:
                stack.append(char)
            elif len(stack) == 0 or pair[stack.pop()] != char:
                return False
        
        return len(stack) == 0
