class Solution:
    def isValid(self, s: str) -> bool:
        open_stack = []
        for c in s:
            if c == "(" or c == "[" or c == "{":
                open_stack.append(c)
            else:
                if len(open_stack) == 0 or not self.matches(open_stack[-1], c):
                    return False
                op = open_stack.pop()
                
        return len(open_stack) == 0
    
    def matches(self, op: str, cl: str) -> bool:
        if op == "(": return cl == ")"
        if op == "[": return cl == "]"
        if op == "{": return cl == "}"
        return False
        