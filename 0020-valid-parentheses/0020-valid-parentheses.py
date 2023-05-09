class Solution:
    def isValid(self, s: str) -> bool:
        open_stack = []
        for c in s:
            if c == "(" or c == "[" or c == "{":
                open_stack.append(c)
            else:
                if len(open_stack) == 0:
                    return False
                op = open_stack.pop()
                if not self.matches(op, c):
                    return False
        return len(open_stack) == 0
    
    def matches(self, op: str, cl: str) -> bool:
        if op == "(": return cl == ")"
        if op == "[": return cl == "]"
        if op == "{": return cl == "}"
        return False
        